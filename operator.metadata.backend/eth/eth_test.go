package eth_test

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"
	"strconv"
	"testing"

	"ssv_operator_metadata/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Test_Verify(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fmt.Println(crypto.PubkeyToAddress(*publicKeyECDSA))

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x2c565d72847326a08648c9307d965c19bfe522940a4011b974d669454877c44b3acf484c2eae02c1f689e198f099bda796b35e1670e03abdb6a8a6931afbfe0900

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}

func Test_Verify1(t *testing.T) {

	message := `===============+++`
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)

	hash := crypto.Keccak256Hash(hashedMessage)

	signature, _ := hexutil.Decode("0x53c598dcfcc83d036b8e54392f9cc9a0c555001ec4cfaf7a076982b5ade6f7f878d3839d106782aeac44864aac0b8528a068d08d7d5b7c10691db0088895dfec1c")

	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if sigPublicKeyECDSA == nil {
		fmt.Println("==sigPublicKeyECDSA====", err)

	}
	fmt.Println("==================", crypto.PubkeyToAddress(*sigPublicKeyECDSA).String())
}

func Test_Verify_Personal_Sign(t *testing.T) {
	message := "Example `personal_sign` message"

	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)

	hash := crypto.Keccak256Hash(hashedMessage)

	signature, _ := hexutil.Decode("0xa5cae1e72ef181f98c90380d99e9b1a56010524ab9a7658a34d0e0768fcb1e0e090234594d9767ee309841aa9276147bf1cc7a147fca6bc7ee557d230ff939801b")

	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if sigPublicKeyECDSA == nil {
		fmt.Println("==sigPublicKeyECDSA====", err)

	}
	fmt.Println("==================", crypto.PubkeyToAddress(*sigPublicKeyECDSA).String())
}

func Test_Verify_Type_Data_Sign_V4(t *testing.T) {

	data := `{"domain":{"chainId":"5","name":"Ether Mail"},"message":{"contents":"Hi bob"},"primaryType":"Mail","types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"chainId","type":"uint256"}],"Mail":[{"name":"contents","type":"string"}]}}`
	signature := "0xc886bd6580b2c2b51f977e6235bc21bf90a0e61b57e07220ed98b358fbf4e38b2deca96c1e16e254da4f10781de4fe816b19fc3a2b8c7347f15dc1b48ddc290b1c"
	addresss := common.HexToAddress("0x2aFa1332ac57b7162CD2dE90882b3A904979fB12")

	pass, err := eth.VerifyTypeDataSignV4(data, signature, addresss)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("pass:", pass)
}

func Test_NewEthClient(t *testing.T) {
	eth.NewEthClient(false, "https://goerli.infura.io/v3/9e8bab40a2bf4e9bb007f8c35d08835d")
}

func Test_Get_Operator_By_Id(t *testing.T) {
	_cli := eth.NewEthClient(false, "https://goerli.infura.io/v3/9e8bab40a2bf4e9bb007f8c35d08835d")
	v1, v2, v3, v4, v5, v6 := _cli.GetOperatorById(604)
	fmt.Printf("v1:%s\nv2:%v\nv3:%v\nv4:%v\nv5:%v\nv6:%v\n", v1.String(), v2, v3, v4, v5, v6)
}

// ssv_network_test

func TestSSVContact(t *testing.T) {
	// rpc := "https://goerli.infura.io/v3/"
	rpc := "https://goerli.infura.io/v3/9e8bab40a2bf4e9bb007f8c35d08835d"
	ssvNetworkAddress := "0x8dB45282d7C4559fd093C26f677B3837a5598914"
	_cli, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		panic(err)
	}
	// fmt.Println(_cli.ChainID(context.Background()))
	// return
	ssvNetCaller, err := eth.NewContract(common.HexToAddress(ssvNetworkAddress), _cli)
	if err != nil {
		fmt.Println("============++++++000++", err)
		panic(err)
	}

	fmt.Println("============++++++2222++", err)

	// struct Operator {
	//     string name;
	//     bytes publicKey;
	//     uint64 fee;
	//     address ownerAddress;
	//     uint32 score;
	//     uint32 indexInOwner;
	//     uint32 validatorCount;
	//     bool active;
	// }
	// return (operator.name, operator.ownerAddress, operator.publicKey, _operators[operatorId].validatorCount, operator.fee.expand(), operator.score, operator.active);

	// name, ownerAddress, publicKey, validatorCount, fee, score, active, err := ssvNetCaller.GetOperatorById(nil, 1)
	operatorOwner, fee, validatorCount, isPrivate, isActive, err := ssvNetCaller.GetOperatorById(nil, 1)
	fmt.Println(operatorOwner, fee, validatorCount, isPrivate, isActive, err)
}
