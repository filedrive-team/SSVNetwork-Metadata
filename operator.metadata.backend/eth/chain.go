package eth

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/ipfs/go-log/v2"
	"go.uber.org/zap"

	"ssv_operator_metadata/db/models"
	"ssv_operator_metadata/utils"
)

var (
	ssvNetInfo = map[string]*ClientConfig{
		// Ethereum Mainnet
		"mainnet": {
			TokenAddr: "0x9D65fF81a3c488d585bBfb0Bfe3c7707c7917f54",
		},
		// Goerli Testnet
		"testnet": {
			TokenAddr:           "0x3a9f01091C446bdE031E39ea8354647AFef091E7",
			NetworkAddr:         "0xAfdb141Dd99b5a101065f40e3D7636262dce65b3",
			NetworkRegistryAddr: "0x8dB45282d7C4559fd093C26f677B3837a5598914",
			SsvRegistryTopic:    "0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4",
			StartBlockHeight:    big.NewInt(8664885),
		},
	}
	// startBlockHeight = big.NewInt(7287230)
	logging = log.Logger("eth")
)

type ClientConfig struct {
	EthUrl              string
	TokenAddr           string
	NetworkAddr         string
	NetworkRegistryAddr string
	SsvRegistryTopic    string
	StartBlockHeight    *big.Int
}
type Client struct {
	cli         *ethclient.Client
	ssvContract *Contract
	cfg         *ClientConfig
	bk          *utils.BackOff
}

func NewEthClient(isMainNet bool, url string) *Client {
	var cfg *ClientConfig
	if isMainNet {
		cfg = ssvNetInfo["mainnet"]
	} else {
		cfg = ssvNetInfo["testnet"]
	}
	cfg.EthUrl = url
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	gClient, err := ethclient.DialContext(ctx, cfg.EthUrl)

	// gClient, err := ethclient.Dial(cfg.EthUrl)
	if err != nil {
		panic(fmt.Errorf("[NewEthClient] ethclient.Dial error:%v", err))
	}
	// check
	chainId, err := gClient.NetworkID(context.Background())
	if err != nil {
		panic(fmt.Errorf("[NewEthClient] NetworkID faild"))
	}
	fmt.Println("eth chain id :", chainId.String())

	ssvContract, err := NewContract(common.HexToAddress(cfg.NetworkAddr), gClient)
	if err != nil {
		panic(fmt.Errorf("[NewEthClient] NewContract error:%v", err))
	}
	bk := utils.NewDefaultBackoff()
	return &Client{
		cli:         gClient,
		ssvContract: ssvContract,
		cfg:         cfg,
		bk:          bk,
	}
}
func (c *Client) GetOperatorById(id uint64) (common.Address, *big.Int, uint32, bool, bool, error) {
	operatorOwner, fee, validatorCount, isPrivate, isActive, err := c.ssvContract.GetOperatorById(&bind.CallOpts{}, id)
	if err != nil {
		return operatorOwner, fee, validatorCount, isPrivate, isActive, err
	}
	return operatorOwner, fee, validatorCount, isPrivate, isActive, nil
}

func (c *Client) filterSSVContract(ctx context.Context, startHeight *big.Int, info chan interface{}) {
	height, _ := c.cli.BlockNumber(ctx)
	fmt.Println("curr chain height:", height)
	if startHeight == nil {
		startHeight = c.cfg.StartBlockHeight
	}
	div := int64(10000)
	for start := startHeight; start.Cmp(big.NewInt(int64(height))) <= 0; {

		logs, err := c.cli.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: new(big.Int).Set(startHeight), //begin
			ToBlock:   startHeight.Add(startHeight, big.NewInt(div)),
			Addresses: []common.Address{common.HexToAddress(c.cfg.NetworkAddr)},
			Topics:    [][]common.Hash{{common.HexToHash(c.cfg.SsvRegistryTopic)}},
		})

		if err != nil {
			logging.Error("IlterLog error:", err)
			// fmt.Println("IlterLog error:", err)
			return
		}
		for _, vLog := range logs {

			contractAbi := ssvNetworkContractABI()
			parsed, err := ParseOperatorAddedEvent(vLog, contractAbi)
			if err != nil {
				logging.Error("=====ParseOperatorAddedEvent\n", err)
			} else {
				logging.Debug("==data:===%v\n", string(parsed.PublicKey))
				logging.Debug("------", parsed.OperatorId)
				fmt.Printf("=============%v\n", parsed.Fee.String())
				// fmt.Printf("==========%v\n%v\n%v\n%v\n", parsed.OwnerAddress.Hex(), string(parsed.PublicKey), parsed.Id, parsed.Name)
				_, _, validatorCount, _, isActive, err := c.ssvContract.GetOperatorById(nil, uint64(parsed.OperatorId))
				// logging.Debug("pub key cmp:", bytes.Compare(publicKey, parsed.PublicKey))
				// logging.Debug("query:", len(publicKey), string(publicKey))
				logging.Debug("create:", len(parsed.PublicKey), string(parsed.PublicKey))
				logging.Debug(parsed.Owner.String(), parsed.Fee, validatorCount, isActive, err)
				info <- models.OperatorBaseInfo{
					Id:             parsed.OperatorId,
					OwnerAddress:   parsed.Owner.String(),
					PublicKey:      string(parsed.PublicKey),
					Active:         isActive,
					Fee:            parsed.Fee.String(),
					ValidatorCount: validatorCount,
					Timestamp:      int64(vLog.BlockNumber),
				}
			}
		}
		time.Sleep(time.Microsecond * 100)
	}
}

func (c *Client) subSSVContractLog(ctx context.Context, info chan interface{}) {
	defer func() {
		if err := recover(); err != nil {
			delayTime := c.bk.NextBackOff()
			logging.Warnf("subSSVContractLog panic : %v \nnext delay : %v", err, delayTime)
			time.Sleep(delayTime)

			logging.Info("subSSVContractLog reconnect and subscribe")
			// reconnect eth client
			c.cli, err = ethclient.DialContext(ctx, c.cfg.EthUrl)
			if err != nil {
				panic(fmt.Errorf("subSSVContractLog recover error :%v", err))
			}
			// sub
			go c.subSSVContractLog(ctx, info)
		}
	}()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(c.cfg.NetworkAddr)},
		Topics:    [][]common.Hash{{common.HexToHash(c.cfg.SsvRegistryTopic)}},
	}
	_logs := make(chan types.Log)
	sub, err := c.cli.SubscribeFilterLogs(ctx, query, _logs)
	if err != nil {
		e := fmt.Errorf("SubscribeFilterLogs error: %v ", err)
		logging.Error(e)
		panic(e)
	}
	// suscribe success , do back off reset
	c.bk.Reset()

	for {
		select {
		case err := <-sub.Err():
			e := fmt.Errorf("subSSVContractLog sub error: %v ", err)
			logging.Info(e)
			panic(e)
		case vLog := <-_logs:
			logging.Debug(vLog)
			if vLog.Removed {
				continue
			}
			contractAbi := ssvNetworkContractABI()
			// abiParser := eth1.NewParser(zap.NewNop(), eth1.V2)
			eventName := ""
			if ev, err := contractAbi.EventByID(vLog.Topics[0]); err == nil && ev != nil {
				eventName = ev.Name
			}

			if err != nil {
				continue
			}

			parsed, err := ParseOperatorAddedEvent(vLog, contractAbi)
			if err != nil {
				logging.Warn("could not parse ongoing event, the event is malformed",
					zap.String("event", eventName),
					zap.Uint64("block", vLog.BlockNumber),
					zap.String("txHash", vLog.TxHash.Hex()),
					zap.Error(err),
				)
				continue
			}

			_, _, validatorCount, _, isActive, err := c.ssvContract.GetOperatorById(nil, parsed.OperatorId)

			if err != nil {
				logging.Warn("could not get operator by id, the event is malformed",
					zap.String("event", eventName),
					zap.Uint64("block", vLog.BlockNumber),
					zap.String("txHash", vLog.TxHash.Hex()),
					zap.Error(err),
				)
				continue
			}

			info <- models.OperatorBaseInfo{
				Id:             parsed.OperatorId,
				OwnerAddress:   parsed.Owner.String(),
				PublicKey:      string(parsed.PublicKey),
				Active:         isActive,
				Fee:            parsed.Fee.String(),
				ValidatorCount: validatorCount,
				Timestamp:      int64(vLog.BlockNumber),
			}

		case <-ctx.Done():
			logging.Info("subSSVContractLog done")
			return
		}
	}
}

func (c *Client) SyncSSVContractInfo(ctx context.Context, isAllFilter bool, chInfo chan interface{}) {
	// 1, filter logs | get last height and filter
	if isAllFilter {
		c.filterSSVContract(ctx, c.cfg.StartBlockHeight, chInfo)
	}
	// 2, Subscribe Filter Logs
	go c.subSSVContractLog(ctx, chInfo)
}

func ssvNetworkContractABI() abi.ABI {
	// contractAbiString := eth1.ContractABI(v)
	contractAbi, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		fmt.Printf("=====%+v", err)
	} else {
		//fmt.Printf("========%+v\n", contractAbi)
		//fmt.Printf("========%+v\n", contractAbiString)
	}
	return contractAbi
}

func ParseOperatorAddedEvent(log types.Log, contractAbi abi.ABI) (*models.OperatorAddedEvent, error) {
	var event models.OperatorAddedEvent
	err := contractAbi.UnpackIntoInterface(&event, "OperatorAdded", log.Data)
	if err != nil {
		return nil, fmt.Errorf("could not unpack event")
	}

	pubKey, err := unpackField(event.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("could not read %s event operator public key OperatorAdded", err)
	}
	event.PublicKey = pubKey

	if len(log.Topics) < 3 {
		return nil, fmt.Errorf("%s event missing topics", "OperatorAdded")
	}
	event.OperatorId = log.Topics[1].Big().Uint64()
	event.Owner = common.HexToAddress(log.Topics[2].Hex())

	return &event, nil
}

func unpackField(fieldBytes []byte) ([]byte, error) {
	outAbi, err := getOutAbi()
	if err != nil {
		return nil, fmt.Errorf("could not define ABI", err)
	}

	outField, err := outAbi.Unpack("method", fieldBytes)
	if err != nil {
		return nil, err
	}

	unpacked, ok := outField[0].([]byte)
	if !ok {
		return nil, fmt.Errorf("could not cast OperatorPublicKey", err)
	}

	return unpacked, nil
}

func getOutAbi() (abi.ABI, error) {
	def := `[{ "name" : "method", "type": "function", "outputs": [{"type": "bytes"}]}]`
	return abi.JSON(strings.NewReader(def))
}
