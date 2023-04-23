// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalNotWithinTimeframe\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterDoesNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterIsLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterNotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedValidatorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsIncreaseLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeIncreaseNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectClusterState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIdsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewBlockPeriodIsBelowMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeDelcared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameFeeChangeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsortedOperatorsList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorOwnedByOtherAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeIncreaseNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterReactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"DeclareOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"ExecuteOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"FeeRecipientAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"LiquidationThresholdPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinimumLiquidationCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"NetworkEarningsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"NetworkFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeCancellationDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeIncreaseLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"OperatorWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"cancelDeclaredOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"clusters\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"balance\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"declareOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"declareOperatorFeePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"executeOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeOperatorFeePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"initialVersion_\",\"type\":\"string\"},{\"internalType\":\"contract IERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"operatorMaxFeeIncrease_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"declareOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executeOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minimumBlocksBeforeLiquidation_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidationCollateral_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBlocksBeforeLiquidation\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumLiquidationCollateral\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"networkFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndexBlockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"operatorFeeChangeRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"approvalBeginTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"approvalEndTime\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorMaxFeeIncrease\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"balance\",\"type\":\"uint64\"}],\"internalType\":\"struct ISSVNetworkCore.Snapshot\",\"name\":\"snapshot\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"operatorsWhitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"reactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"reduceOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"setFeeRecipientAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"setOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDeclareOperatorFeePeriod\",\"type\":\"uint64\"}],\"name\":\"updateDeclareOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newExecuteOperatorFeePeriod\",\"type\":\"uint64\"}],\"name\":\"updateExecuteOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blocks\",\"type\":\"uint64\"}],\"name\":\"updateLiquidationThresholdPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinimumLiquidationCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateNetworkFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newOperatorMaxFeeIncrease\",\"type\":\"uint64\"}],\"name\":\"updateOperatorFeeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validatorPKs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorsPerOperatorLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"struct ISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNetworkEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AddressNetworkFee is a free data retrieval call binding the contract method 0x56d9a2cc.
//
// Solidity: function addressNetworkFee(address ownerAddress) view returns(uint256)
func (_Contract *ContractCaller) AddressNetworkFee(opts *bind.CallOpts, ownerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addressNetworkFee", ownerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressNetworkFee is a free data retrieval call binding the contract method 0x56d9a2cc.
//
// Solidity: function addressNetworkFee(address ownerAddress) view returns(uint256)
func (_Contract *ContractSession) AddressNetworkFee(ownerAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.AddressNetworkFee(&_Contract.CallOpts, ownerAddress)
}

// AddressNetworkFee is a free data retrieval call binding the contract method 0x56d9a2cc.
//
// Solidity: function addressNetworkFee(address ownerAddress) view returns(uint256)
func (_Contract *ContractCallerSession) AddressNetworkFee(ownerAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.AddressNetworkFee(&_Contract.CallOpts, ownerAddress)
}

// GetAddressBalance is a free data retrieval call binding the contract method 0x35046722.
//
// Solidity: function getAddressBalance(address ownerAddress) view returns(uint256)
func (_Contract *ContractCaller) GetAddressBalance(opts *bind.CallOpts, ownerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddressBalance", ownerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressBalance is a free data retrieval call binding the contract method 0x35046722.
//
// Solidity: function getAddressBalance(address ownerAddress) view returns(uint256)
func (_Contract *ContractSession) GetAddressBalance(ownerAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.GetAddressBalance(&_Contract.CallOpts, ownerAddress)
}

// GetAddressBalance is a free data retrieval call binding the contract method 0x35046722.
//
// Solidity: function getAddressBalance(address ownerAddress) view returns(uint256)
func (_Contract *ContractCallerSession) GetAddressBalance(ownerAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.GetAddressBalance(&_Contract.CallOpts, ownerAddress)
}

// GetAddressBurnRate is a free data retrieval call binding the contract method 0x2563c64e.
//
// Solidity: function getAddressBurnRate(address ownerAddress) view returns(uint256)
func (_Contract *ContractCaller) GetAddressBurnRate(opts *bind.CallOpts, ownerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddressBurnRate", ownerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressBurnRate is a free data retrieval call binding the contract method 0x2563c64e.
//
// Solidity: function getAddressBurnRate(address ownerAddress) view returns(uint256)
func (_Contract *ContractSession) GetAddressBurnRate(ownerAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.GetAddressBurnRate(&_Contract.CallOpts, ownerAddress)
}

// GetAddressBurnRate is a free data retrieval call binding the contract method 0x2563c64e.
//
// Solidity: function getAddressBurnRate(address ownerAddress) view returns(uint256)
func (_Contract *ContractCallerSession) GetAddressBurnRate(ownerAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.GetAddressBurnRate(&_Contract.CallOpts, ownerAddress)
}

// GetDeclaredOperatorFeePeriod is a free data retrieval call binding the contract method 0x1be2bd74.
//
// Solidity: function getDeclaredOperatorFeePeriod() view returns(uint256)
func (_Contract *ContractCaller) GetDeclaredOperatorFeePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDeclaredOperatorFeePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeclaredOperatorFeePeriod is a free data retrieval call binding the contract method 0x1be2bd74.
//
// Solidity: function getDeclaredOperatorFeePeriod() view returns(uint256)
func (_Contract *ContractSession) GetDeclaredOperatorFeePeriod() (*big.Int, error) {
	return _Contract.Contract.GetDeclaredOperatorFeePeriod(&_Contract.CallOpts)
}

// GetDeclaredOperatorFeePeriod is a free data retrieval call binding the contract method 0x1be2bd74.
//
// Solidity: function getDeclaredOperatorFeePeriod() view returns(uint256)
func (_Contract *ContractCallerSession) GetDeclaredOperatorFeePeriod() (*big.Int, error) {
	return _Contract.Contract.GetDeclaredOperatorFeePeriod(&_Contract.CallOpts)
}

// GetExecuteOperatorFeePeriod is a free data retrieval call binding the contract method 0xdd83fcb6.
//
// Solidity: function getExecuteOperatorFeePeriod() view returns(uint256)
func (_Contract *ContractCaller) GetExecuteOperatorFeePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getExecuteOperatorFeePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecuteOperatorFeePeriod is a free data retrieval call binding the contract method 0xdd83fcb6.
//
// Solidity: function getExecuteOperatorFeePeriod() view returns(uint256)
func (_Contract *ContractSession) GetExecuteOperatorFeePeriod() (*big.Int, error) {
	return _Contract.Contract.GetExecuteOperatorFeePeriod(&_Contract.CallOpts)
}

// GetExecuteOperatorFeePeriod is a free data retrieval call binding the contract method 0xdd83fcb6.
//
// Solidity: function getExecuteOperatorFeePeriod() view returns(uint256)
func (_Contract *ContractCallerSession) GetExecuteOperatorFeePeriod() (*big.Int, error) {
	return _Contract.Contract.GetExecuteOperatorFeePeriod(&_Contract.CallOpts)
}

// GetLiquidationThresholdPeriod is a free data retrieval call binding the contract method 0x9040f7c3.
//
// Solidity: function getLiquidationThresholdPeriod() view returns(uint256)
func (_Contract *ContractCaller) GetLiquidationThresholdPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLiquidationThresholdPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationThresholdPeriod is a free data retrieval call binding the contract method 0x9040f7c3.
//
// Solidity: function getLiquidationThresholdPeriod() view returns(uint256)
func (_Contract *ContractSession) GetLiquidationThresholdPeriod() (*big.Int, error) {
	return _Contract.Contract.GetLiquidationThresholdPeriod(&_Contract.CallOpts)
}

// GetLiquidationThresholdPeriod is a free data retrieval call binding the contract method 0x9040f7c3.
//
// Solidity: function getLiquidationThresholdPeriod() view returns(uint256)
func (_Contract *ContractCallerSession) GetLiquidationThresholdPeriod() (*big.Int, error) {
	return _Contract.Contract.GetLiquidationThresholdPeriod(&_Contract.CallOpts)
}

// GetNetworkEarnings is a free data retrieval call binding the contract method 0x777915cb.
//
// Solidity: function getNetworkEarnings() view returns(uint256)
func (_Contract *ContractCaller) GetNetworkEarnings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNetworkEarnings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkEarnings is a free data retrieval call binding the contract method 0x777915cb.
//
// Solidity: function getNetworkEarnings() view returns(uint256)
func (_Contract *ContractSession) GetNetworkEarnings() (*big.Int, error) {
	return _Contract.Contract.GetNetworkEarnings(&_Contract.CallOpts)
}

// GetNetworkEarnings is a free data retrieval call binding the contract method 0x777915cb.
//
// Solidity: function getNetworkEarnings() view returns(uint256)
func (_Contract *ContractCallerSession) GetNetworkEarnings() (*big.Int, error) {
	return _Contract.Contract.GetNetworkEarnings(&_Contract.CallOpts)
}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_Contract *ContractCaller) GetNetworkFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNetworkFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_Contract *ContractSession) GetNetworkFee() (*big.Int, error) {
	return _Contract.Contract.GetNetworkFee(&_Contract.CallOpts)
}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_Contract *ContractCallerSession) GetNetworkFee() (*big.Int, error) {
	return _Contract.Contract.GetNetworkFee(&_Contract.CallOpts)
}

// GetOperatorById is a free data retrieval call binding the contract method 0x0260d443.
//
// Solidity:     function getOperatorById(uint64 operatorId) external view override returns (address, uint256, uint32, bool, bool) 
// operatorOwner, fee.expand(), validatorCount, isPrivate, isActive
func (_Contract *ContractCaller) GetOperatorById(opts *bind.CallOpts, operatorId uint64) (common.Address, *big.Int, uint32, bool, bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorById", operatorId)
	if err != nil {
		return  *new(common.Address),  *new(*big.Int), *new(uint32),*new(bool), *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)
	return out0, out1, out2, out3, out4, err
}

// GetOperatorById is a free data retrieval call binding the contract method 0x0260d443.
//
// Solidity:     function getOperatorById(uint64 operatorId) external view override returns (address, uint256, uint32, bool, bool) 
// operatorOwner, fee.expand(), validatorCount, isPrivate, isActive
func (_Contract *ContractSession) GetOperatorById(operatorId uint64) ( common.Address, *big.Int, uint32, bool, bool, error) {
	return _Contract.Contract.GetOperatorById(&_Contract.CallOpts, operatorId)
}

// GetOperatorById is a free data retrieval call binding the contract method 0x0260d443.
//
// Solidity:     function getOperatorById(uint64 operatorId) external view override returns (address, uint256, uint32, bool, bool) 
// operatorOwner, fee.expand(), validatorCount, isPrivate, isActive
func (_Contract *ContractCallerSession) GetOperatorById(operatorId uint64)  ( common.Address, *big.Int, uint32, bool, bool, error) {
	return _Contract.Contract.GetOperatorById(&_Contract.CallOpts, operatorId)
}

// GetOperatorByPublicKey is a free data retrieval call binding the contract method 0xd26af797.
//
// Solidity: function getOperatorByPublicKey(bytes publicKey) view returns(string, address, bytes, uint256, uint256, uint256, bool)
func (_Contract *ContractCaller) GetOperatorByPublicKey(opts *bind.CallOpts, publicKey []byte) (string, common.Address, []byte, *big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorByPublicKey", publicKey)

	if err != nil {
		return *new(string), *new(common.Address), *new([]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetOperatorByPublicKey is a free data retrieval call binding the contract method 0xd26af797.
//
// Solidity: function getOperatorByPublicKey(bytes publicKey) view returns(string, address, bytes, uint256, uint256, uint256, bool)
func (_Contract *ContractSession) GetOperatorByPublicKey(publicKey []byte) (string, common.Address, []byte, *big.Int, *big.Int, *big.Int, bool, error) {
	return _Contract.Contract.GetOperatorByPublicKey(&_Contract.CallOpts, publicKey)
}

// GetOperatorByPublicKey is a free data retrieval call binding the contract method 0xd26af797.
//
// Solidity: function getOperatorByPublicKey(bytes publicKey) view returns(string, address, bytes, uint256, uint256, uint256, bool)
func (_Contract *ContractCallerSession) GetOperatorByPublicKey(publicKey []byte) (string, common.Address, []byte, *big.Int, *big.Int, *big.Int, bool, error) {
	return _Contract.Contract.GetOperatorByPublicKey(&_Contract.CallOpts, publicKey)
}

// GetOperatorDeclaredFee is a free data retrieval call binding the contract method 0x85747566.
//
// Solidity: function getOperatorDeclaredFee(uint32 operatorId) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) GetOperatorDeclaredFee(opts *bind.CallOpts, operatorId uint32) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorDeclaredFee", operatorId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetOperatorDeclaredFee is a free data retrieval call binding the contract method 0x85747566.
//
// Solidity: function getOperatorDeclaredFee(uint32 operatorId) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) GetOperatorDeclaredFee(operatorId uint32) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetOperatorDeclaredFee(&_Contract.CallOpts, operatorId)
}

// GetOperatorDeclaredFee is a free data retrieval call binding the contract method 0x85747566.
//
// Solidity: function getOperatorDeclaredFee(uint32 operatorId) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) GetOperatorDeclaredFee(operatorId uint32) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetOperatorDeclaredFee(&_Contract.CallOpts, operatorId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0xf8f08d7f.
//
// Solidity: function getOperatorFee(uint32 operatorId) view returns(uint256)
func (_Contract *ContractCaller) GetOperatorFee(opts *bind.CallOpts, operatorId uint32) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorFee", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorFee is a free data retrieval call binding the contract method 0xf8f08d7f.
//
// Solidity: function getOperatorFee(uint32 operatorId) view returns(uint256)
func (_Contract *ContractSession) GetOperatorFee(operatorId uint32) (*big.Int, error) {
	return _Contract.Contract.GetOperatorFee(&_Contract.CallOpts, operatorId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0xf8f08d7f.
//
// Solidity: function getOperatorFee(uint32 operatorId) view returns(uint256)
func (_Contract *ContractCallerSession) GetOperatorFee(operatorId uint32) (*big.Int, error) {
	return _Contract.Contract.GetOperatorFee(&_Contract.CallOpts, operatorId)
}

// GetOperatorFeeIncreaseLimit is a free data retrieval call binding the contract method 0x68465f7d.
//
// Solidity: function getOperatorFeeIncreaseLimit() view returns(uint256)
func (_Contract *ContractCaller) GetOperatorFeeIncreaseLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorFeeIncreaseLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorFeeIncreaseLimit is a free data retrieval call binding the contract method 0x68465f7d.
//
// Solidity: function getOperatorFeeIncreaseLimit() view returns(uint256)
func (_Contract *ContractSession) GetOperatorFeeIncreaseLimit() (*big.Int, error) {
	return _Contract.Contract.GetOperatorFeeIncreaseLimit(&_Contract.CallOpts)
}

// GetOperatorFeeIncreaseLimit is a free data retrieval call binding the contract method 0x68465f7d.
//
// Solidity: function getOperatorFeeIncreaseLimit() view returns(uint256)
func (_Contract *ContractCallerSession) GetOperatorFeeIncreaseLimit() (*big.Int, error) {
	return _Contract.Contract.GetOperatorFeeIncreaseLimit(&_Contract.CallOpts)
}

// GetOperatorsByValidator is a free data retrieval call binding the contract method 0x053e8349.
//
// Solidity: function getOperatorsByValidator(bytes publicKey) view returns(uint32[])
func (_Contract *ContractCaller) GetOperatorsByValidator(opts *bind.CallOpts, publicKey []byte) ([]uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorsByValidator", publicKey)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetOperatorsByValidator is a free data retrieval call binding the contract method 0x053e8349.
//
// Solidity: function getOperatorsByValidator(bytes publicKey) view returns(uint32[])
func (_Contract *ContractSession) GetOperatorsByValidator(publicKey []byte) ([]uint32, error) {
	return _Contract.Contract.GetOperatorsByValidator(&_Contract.CallOpts, publicKey)
}

// GetOperatorsByValidator is a free data retrieval call binding the contract method 0x053e8349.
//
// Solidity: function getOperatorsByValidator(bytes publicKey) view returns(uint32[])
func (_Contract *ContractCallerSession) GetOperatorsByValidator(publicKey []byte) ([]uint32, error) {
	return _Contract.Contract.GetOperatorsByValidator(&_Contract.CallOpts, publicKey)
}

// GetValidatorsByOwnerAddress is a free data retrieval call binding the contract method 0x57c7ce22.
//
// Solidity: function getValidatorsByOwnerAddress(address ownerAddress) view returns(bytes[])
func (_Contract *ContractCaller) GetValidatorsByOwnerAddress(opts *bind.CallOpts, ownerAddress common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorsByOwnerAddress", ownerAddress)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetValidatorsByOwnerAddress is a free data retrieval call binding the contract method 0x57c7ce22.
//
// Solidity: function getValidatorsByOwnerAddress(address ownerAddress) view returns(bytes[])
func (_Contract *ContractSession) GetValidatorsByOwnerAddress(ownerAddress common.Address) ([][]byte, error) {
	return _Contract.Contract.GetValidatorsByOwnerAddress(&_Contract.CallOpts, ownerAddress)
}

// GetValidatorsByOwnerAddress is a free data retrieval call binding the contract method 0x57c7ce22.
//
// Solidity: function getValidatorsByOwnerAddress(address ownerAddress) view returns(bytes[])
func (_Contract *ContractCallerSession) GetValidatorsByOwnerAddress(ownerAddress common.Address) ([][]byte, error) {
	return _Contract.Contract.GetValidatorsByOwnerAddress(&_Contract.CallOpts, ownerAddress)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address ownerAddress) view returns(bool)
func (_Contract *ContractCaller) IsLiquidatable(opts *bind.CallOpts, ownerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isLiquidatable", ownerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address ownerAddress) view returns(bool)
func (_Contract *ContractSession) IsLiquidatable(ownerAddress common.Address) (bool, error) {
	return _Contract.Contract.IsLiquidatable(&_Contract.CallOpts, ownerAddress)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address ownerAddress) view returns(bool)
func (_Contract *ContractCallerSession) IsLiquidatable(ownerAddress common.Address) (bool, error) {
	return _Contract.Contract.IsLiquidatable(&_Contract.CallOpts, ownerAddress)
}

// IsLiquidated is a free data retrieval call binding the contract method 0xa3b3f606.
//
// Solidity: function isLiquidated(address ownerAddress) view returns(bool)
func (_Contract *ContractCaller) IsLiquidated(opts *bind.CallOpts, ownerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isLiquidated", ownerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidated is a free data retrieval call binding the contract method 0xa3b3f606.
//
// Solidity: function isLiquidated(address ownerAddress) view returns(bool)
func (_Contract *ContractSession) IsLiquidated(ownerAddress common.Address) (bool, error) {
	return _Contract.Contract.IsLiquidated(&_Contract.CallOpts, ownerAddress)
}

// IsLiquidated is a free data retrieval call binding the contract method 0xa3b3f606.
//
// Solidity: function isLiquidated(address ownerAddress) view returns(bool)
func (_Contract *ContractCallerSession) IsLiquidated(ownerAddress common.Address) (bool, error) {
	return _Contract.Contract.IsLiquidated(&_Contract.CallOpts, ownerAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ValidatorsPerOperatorCount is a free data retrieval call binding the contract method 0xd379a5f2.
//
// Solidity: function validatorsPerOperatorCount(uint32 operatorId) view returns(uint32)
func (_Contract *ContractCaller) ValidatorsPerOperatorCount(opts *bind.CallOpts, operatorId uint32) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorsPerOperatorCount", operatorId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ValidatorsPerOperatorCount is a free data retrieval call binding the contract method 0xd379a5f2.
//
// Solidity: function validatorsPerOperatorCount(uint32 operatorId) view returns(uint32)
func (_Contract *ContractSession) ValidatorsPerOperatorCount(operatorId uint32) (uint32, error) {
	return _Contract.Contract.ValidatorsPerOperatorCount(&_Contract.CallOpts, operatorId)
}

// ValidatorsPerOperatorCount is a free data retrieval call binding the contract method 0xd379a5f2.
//
// Solidity: function validatorsPerOperatorCount(uint32 operatorId) view returns(uint32)
func (_Contract *ContractCallerSession) ValidatorsPerOperatorCount(operatorId uint32) (uint32, error) {
	return _Contract.Contract.ValidatorsPerOperatorCount(&_Contract.CallOpts, operatorId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint32)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint32)
func (_Contract *ContractSession) Version() (uint32, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint32)
func (_Contract *ContractCallerSession) Version() (uint32, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x154996bb.
//
// Solidity: function cancelDeclaredOperatorFee(uint32 operatorId) returns()
func (_Contract *ContractTransactor) CancelDeclaredOperatorFee(opts *bind.TransactOpts, operatorId uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelDeclaredOperatorFee", operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x154996bb.
//
// Solidity: function cancelDeclaredOperatorFee(uint32 operatorId) returns()
func (_Contract *ContractSession) CancelDeclaredOperatorFee(operatorId uint32) (*types.Transaction, error) {
	return _Contract.Contract.CancelDeclaredOperatorFee(&_Contract.TransactOpts, operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x154996bb.
//
// Solidity: function cancelDeclaredOperatorFee(uint32 operatorId) returns()
func (_Contract *ContractTransactorSession) CancelDeclaredOperatorFee(operatorId uint32) (*types.Transaction, error) {
	return _Contract.Contract.CancelDeclaredOperatorFee(&_Contract.TransactOpts, operatorId)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0x70f97314.
//
// Solidity: function declareOperatorFee(uint32 operatorId, uint256 operatorFee) returns()
func (_Contract *ContractTransactor) DeclareOperatorFee(opts *bind.TransactOpts, operatorId uint32, operatorFee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "declareOperatorFee", operatorId, operatorFee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0x70f97314.
//
// Solidity: function declareOperatorFee(uint32 operatorId, uint256 operatorFee) returns()
func (_Contract *ContractSession) DeclareOperatorFee(operatorId uint32, operatorFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeclareOperatorFee(&_Contract.TransactOpts, operatorId, operatorFee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0x70f97314.
//
// Solidity: function declareOperatorFee(uint32 operatorId, uint256 operatorFee) returns()
func (_Contract *ContractTransactorSession) DeclareOperatorFee(operatorId uint32, operatorFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeclareOperatorFee(&_Contract.TransactOpts, operatorId, operatorFee)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address ownerAddress, uint256 amount) returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, ownerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", ownerAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address ownerAddress, uint256 amount) returns()
func (_Contract *ContractSession) Deposit(ownerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, ownerAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address ownerAddress, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Deposit(ownerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, ownerAddress, amount)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x400aa7ce.
//
// Solidity: function executeOperatorFee(uint32 operatorId) returns()
func (_Contract *ContractTransactor) ExecuteOperatorFee(opts *bind.TransactOpts, operatorId uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executeOperatorFee", operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x400aa7ce.
//
// Solidity: function executeOperatorFee(uint32 operatorId) returns()
func (_Contract *ContractSession) ExecuteOperatorFee(operatorId uint32) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteOperatorFee(&_Contract.TransactOpts, operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x400aa7ce.
//
// Solidity: function executeOperatorFee(uint32 operatorId) returns()
func (_Contract *ContractTransactorSession) ExecuteOperatorFee(operatorId uint32) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteOperatorFee(&_Contract.TransactOpts, operatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0x36881f18.
//
// Solidity: function initialize(address registryAddress_, address token_, uint64 minimumBlocksBeforeLiquidation_, uint64 operatorMaxFeeIncrease_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, registryAddress_ common.Address, token_ common.Address, minimumBlocksBeforeLiquidation_ uint64, operatorMaxFeeIncrease_ uint64, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", registryAddress_, token_, minimumBlocksBeforeLiquidation_, operatorMaxFeeIncrease_, declareOperatorFeePeriod_, executeOperatorFeePeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0x36881f18.
//
// Solidity: function initialize(address registryAddress_, address token_, uint64 minimumBlocksBeforeLiquidation_, uint64 operatorMaxFeeIncrease_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_) returns()
func (_Contract *ContractSession) Initialize(registryAddress_ common.Address, token_ common.Address, minimumBlocksBeforeLiquidation_ uint64, operatorMaxFeeIncrease_ uint64, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, registryAddress_, token_, minimumBlocksBeforeLiquidation_, operatorMaxFeeIncrease_, declareOperatorFeePeriod_, executeOperatorFeePeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0x36881f18.
//
// Solidity: function initialize(address registryAddress_, address token_, uint64 minimumBlocksBeforeLiquidation_, uint64 operatorMaxFeeIncrease_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_) returns()
func (_Contract *ContractTransactorSession) Initialize(registryAddress_ common.Address, token_ common.Address, minimumBlocksBeforeLiquidation_ uint64, operatorMaxFeeIncrease_ uint64, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, registryAddress_, token_, minimumBlocksBeforeLiquidation_, operatorMaxFeeIncrease_, declareOperatorFeePeriod_, executeOperatorFeePeriod_)
}

// Liquidate is a paid mutator transaction binding the contract method 0xa985994b.
//
// Solidity: function liquidate(address[] ownerAddresses) returns()
func (_Contract *ContractTransactor) Liquidate(opts *bind.TransactOpts, ownerAddresses []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "liquidate", ownerAddresses)
}

// Liquidate is a paid mutator transaction binding the contract method 0xa985994b.
//
// Solidity: function liquidate(address[] ownerAddresses) returns()
func (_Contract *ContractSession) Liquidate(ownerAddresses []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Liquidate(&_Contract.TransactOpts, ownerAddresses)
}

// Liquidate is a paid mutator transaction binding the contract method 0xa985994b.
//
// Solidity: function liquidate(address[] ownerAddresses) returns()
func (_Contract *ContractTransactorSession) Liquidate(ownerAddresses []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Liquidate(&_Contract.TransactOpts, ownerAddresses)
}

// ReactivateAccount is a paid mutator transaction binding the contract method 0x2afe872f.
//
// Solidity: function reactivateAccount(uint256 amount) returns()
func (_Contract *ContractTransactor) ReactivateAccount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reactivateAccount", amount)
}

// ReactivateAccount is a paid mutator transaction binding the contract method 0x2afe872f.
//
// Solidity: function reactivateAccount(uint256 amount) returns()
func (_Contract *ContractSession) ReactivateAccount(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReactivateAccount(&_Contract.TransactOpts, amount)
}

// ReactivateAccount is a paid mutator transaction binding the contract method 0x2afe872f.
//
// Solidity: function reactivateAccount(uint256 amount) returns()
func (_Contract *ContractTransactorSession) ReactivateAccount(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReactivateAccount(&_Contract.TransactOpts, amount)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xb083ab35.
//
// Solidity: function registerOperator(string name, bytes publicKey, uint256 fee) returns(uint32 operatorId)
func (_Contract *ContractTransactor) RegisterOperator(opts *bind.TransactOpts, name string, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerOperator", name, publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xb083ab35.
//
// Solidity: function registerOperator(string name, bytes publicKey, uint256 fee) returns(uint32 operatorId)
func (_Contract *ContractSession) RegisterOperator(name string, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterOperator(&_Contract.TransactOpts, name, publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xb083ab35.
//
// Solidity: function registerOperator(string name, bytes publicKey, uint256 fee) returns(uint32 operatorId)
func (_Contract *ContractTransactorSession) RegisterOperator(name string, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterOperator(&_Contract.TransactOpts, name, publicKey, fee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe40cb19d.
//
// Solidity: function registerValidator(bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] sharesEncrypted, uint256 amount) returns()
func (_Contract *ContractTransactor) RegisterValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint32, sharesPublicKeys [][]byte, sharesEncrypted [][]byte, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerValidator", publicKey, operatorIds, sharesPublicKeys, sharesEncrypted, amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe40cb19d.
//
// Solidity: function registerValidator(bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] sharesEncrypted, uint256 amount) returns()
func (_Contract *ContractSession) RegisterValidator(publicKey []byte, operatorIds []uint32, sharesPublicKeys [][]byte, sharesEncrypted [][]byte, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterValidator(&_Contract.TransactOpts, publicKey, operatorIds, sharesPublicKeys, sharesEncrypted, amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe40cb19d.
//
// Solidity: function registerValidator(bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] sharesEncrypted, uint256 amount) returns()
func (_Contract *ContractTransactorSession) RegisterValidator(publicKey []byte, operatorIds []uint32, sharesPublicKeys [][]byte, sharesEncrypted [][]byte, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterValidator(&_Contract.TransactOpts, publicKey, operatorIds, sharesPublicKeys, sharesEncrypted, amount)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e1d2f05.
//
// Solidity: function removeOperator(uint32 operatorId) returns()
func (_Contract *ContractTransactor) RemoveOperator(opts *bind.TransactOpts, operatorId uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeOperator", operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e1d2f05.
//
// Solidity: function removeOperator(uint32 operatorId) returns()
func (_Contract *ContractSession) RemoveOperator(operatorId uint32) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e1d2f05.
//
// Solidity: function removeOperator(uint32 operatorId) returns()
func (_Contract *ContractTransactorSession) RemoveOperator(operatorId uint32) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, operatorId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xb2f569c5.
//
// Solidity: function removeValidator(bytes publicKey) returns()
func (_Contract *ContractTransactor) RemoveValidator(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeValidator", publicKey)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xb2f569c5.
//
// Solidity: function removeValidator(bytes publicKey) returns()
func (_Contract *ContractSession) RemoveValidator(publicKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.RemoveValidator(&_Contract.TransactOpts, publicKey)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xb2f569c5.
//
// Solidity: function removeValidator(bytes publicKey) returns()
func (_Contract *ContractTransactorSession) RemoveValidator(publicKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.RemoveValidator(&_Contract.TransactOpts, publicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 newDeclareOperatorFeePeriod) returns()
func (_Contract *ContractTransactor) UpdateDeclareOperatorFeePeriod(opts *bind.TransactOpts, newDeclareOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateDeclareOperatorFeePeriod", newDeclareOperatorFeePeriod)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 newDeclareOperatorFeePeriod) returns()
func (_Contract *ContractSession) UpdateDeclareOperatorFeePeriod(newDeclareOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDeclareOperatorFeePeriod(&_Contract.TransactOpts, newDeclareOperatorFeePeriod)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 newDeclareOperatorFeePeriod) returns()
func (_Contract *ContractTransactorSession) UpdateDeclareOperatorFeePeriod(newDeclareOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDeclareOperatorFeePeriod(&_Contract.TransactOpts, newDeclareOperatorFeePeriod)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 newExecuteOperatorFeePeriod) returns()
func (_Contract *ContractTransactor) UpdateExecuteOperatorFeePeriod(opts *bind.TransactOpts, newExecuteOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateExecuteOperatorFeePeriod", newExecuteOperatorFeePeriod)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 newExecuteOperatorFeePeriod) returns()
func (_Contract *ContractSession) UpdateExecuteOperatorFeePeriod(newExecuteOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateExecuteOperatorFeePeriod(&_Contract.TransactOpts, newExecuteOperatorFeePeriod)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 newExecuteOperatorFeePeriod) returns()
func (_Contract *ContractTransactorSession) UpdateExecuteOperatorFeePeriod(newExecuteOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateExecuteOperatorFeePeriod(&_Contract.TransactOpts, newExecuteOperatorFeePeriod)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Contract *ContractTransactor) UpdateLiquidationThresholdPeriod(opts *bind.TransactOpts, blocks uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateLiquidationThresholdPeriod", blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Contract *ContractSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateLiquidationThresholdPeriod(&_Contract.TransactOpts, blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Contract *ContractTransactorSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateLiquidationThresholdPeriod(&_Contract.TransactOpts, blocks)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Contract *ContractTransactor) UpdateNetworkFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkFee", fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Contract *ContractSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkFee(&_Contract.TransactOpts, fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkFee(&_Contract.TransactOpts, fee)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 newOperatorMaxFeeIncrease) returns()
func (_Contract *ContractTransactor) UpdateOperatorFeeIncreaseLimit(opts *bind.TransactOpts, newOperatorMaxFeeIncrease uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateOperatorFeeIncreaseLimit", newOperatorMaxFeeIncrease)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 newOperatorMaxFeeIncrease) returns()
func (_Contract *ContractSession) UpdateOperatorFeeIncreaseLimit(newOperatorMaxFeeIncrease uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOperatorFeeIncreaseLimit(&_Contract.TransactOpts, newOperatorMaxFeeIncrease)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 newOperatorMaxFeeIncrease) returns()
func (_Contract *ContractTransactorSession) UpdateOperatorFeeIncreaseLimit(newOperatorMaxFeeIncrease uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOperatorFeeIncreaseLimit(&_Contract.TransactOpts, newOperatorMaxFeeIncrease)
}

// UpdateOperatorScore is a paid mutator transaction binding the contract method 0x5f10ac30.
//
// Solidity: function updateOperatorScore(uint32 operatorId, uint32 score) returns()
func (_Contract *ContractTransactor) UpdateOperatorScore(opts *bind.TransactOpts, operatorId uint32, score uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateOperatorScore", operatorId, score)
}

// UpdateOperatorScore is a paid mutator transaction binding the contract method 0x5f10ac30.
//
// Solidity: function updateOperatorScore(uint32 operatorId, uint32 score) returns()
func (_Contract *ContractSession) UpdateOperatorScore(operatorId uint32, score uint32) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOperatorScore(&_Contract.TransactOpts, operatorId, score)
}

// UpdateOperatorScore is a paid mutator transaction binding the contract method 0x5f10ac30.
//
// Solidity: function updateOperatorScore(uint32 operatorId, uint32 score) returns()
func (_Contract *ContractTransactorSession) UpdateOperatorScore(operatorId uint32, score uint32) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOperatorScore(&_Contract.TransactOpts, operatorId, score)
}

// UpdateValidator is a paid mutator transaction binding the contract method 0x21f7957b.
//
// Solidity: function updateValidator(bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] sharesEncrypted, uint256 amount) returns()
func (_Contract *ContractTransactor) UpdateValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint32, sharesPublicKeys [][]byte, sharesEncrypted [][]byte, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidator", publicKey, operatorIds, sharesPublicKeys, sharesEncrypted, amount)
}

// UpdateValidator is a paid mutator transaction binding the contract method 0x21f7957b.
//
// Solidity: function updateValidator(bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] sharesEncrypted, uint256 amount) returns()
func (_Contract *ContractSession) UpdateValidator(publicKey []byte, operatorIds []uint32, sharesPublicKeys [][]byte, sharesEncrypted [][]byte, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidator(&_Contract.TransactOpts, publicKey, operatorIds, sharesPublicKeys, sharesEncrypted, amount)
}

// UpdateValidator is a paid mutator transaction binding the contract method 0x21f7957b.
//
// Solidity: function updateValidator(bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] sharesEncrypted, uint256 amount) returns()
func (_Contract *ContractTransactorSession) UpdateValidator(publicKey []byte, operatorIds []uint32, sharesPublicKeys [][]byte, sharesEncrypted [][]byte, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidator(&_Contract.TransactOpts, publicKey, operatorIds, sharesPublicKeys, sharesEncrypted, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Contract *ContractSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Contract *ContractTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Contract *ContractSession) WithdrawAll() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAll(&_Contract.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Contract *ContractTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAll(&_Contract.TransactOpts)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Contract *ContractTransactor) WithdrawNetworkEarnings(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawNetworkEarnings", amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Contract *ContractSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawNetworkEarnings(&_Contract.TransactOpts, amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Contract *ContractTransactorSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawNetworkEarnings(&_Contract.TransactOpts, amount)
}

// ContractAccountEnableIterator is returned from FilterAccountEnable and is used to iterate over the raw logs and unpacked data for AccountEnable events raised by the Contract contract.
type ContractAccountEnableIterator struct {
	Event *ContractAccountEnable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAccountEnableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAccountEnable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAccountEnable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAccountEnableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAccountEnableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAccountEnable represents a AccountEnable event raised by the Contract contract.
type ContractAccountEnable struct {
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountEnable is a free log retrieval operation binding the contract event 0xd2cad6abfa380a10a84c0e561544f8deb6eea7dd29a411fa306f902aa6ac5eca.
//
// Solidity: event AccountEnable(address indexed ownerAddress)
func (_Contract *ContractFilterer) FilterAccountEnable(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractAccountEnableIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AccountEnable", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractAccountEnableIterator{contract: _Contract.contract, event: "AccountEnable", logs: logs, sub: sub}, nil
}

// WatchAccountEnable is a free log subscription operation binding the contract event 0xd2cad6abfa380a10a84c0e561544f8deb6eea7dd29a411fa306f902aa6ac5eca.
//
// Solidity: event AccountEnable(address indexed ownerAddress)
func (_Contract *ContractFilterer) WatchAccountEnable(opts *bind.WatchOpts, sink chan<- *ContractAccountEnable, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AccountEnable", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAccountEnable)
				if err := _Contract.contract.UnpackLog(event, "AccountEnable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountEnable is a log parse operation binding the contract event 0xd2cad6abfa380a10a84c0e561544f8deb6eea7dd29a411fa306f902aa6ac5eca.
//
// Solidity: event AccountEnable(address indexed ownerAddress)
func (_Contract *ContractFilterer) ParseAccountEnable(log types.Log) (*ContractAccountEnable, error) {
	event := new(ContractAccountEnable)
	if err := _Contract.contract.UnpackLog(event, "AccountEnable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAccountLiquidationIterator is returned from FilterAccountLiquidation and is used to iterate over the raw logs and unpacked data for AccountLiquidation events raised by the Contract contract.
type ContractAccountLiquidationIterator struct {
	Event *ContractAccountLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAccountLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAccountLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAccountLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAccountLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAccountLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAccountLiquidation represents a AccountLiquidation event raised by the Contract contract.
type ContractAccountLiquidation struct {
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountLiquidation is a free log retrieval operation binding the contract event 0x4fa857769e6aa8d98c53ea164735f03821828e6672f227ea817ec1e74d61c343.
//
// Solidity: event AccountLiquidation(address indexed ownerAddress)
func (_Contract *ContractFilterer) FilterAccountLiquidation(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractAccountLiquidationIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AccountLiquidation", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractAccountLiquidationIterator{contract: _Contract.contract, event: "AccountLiquidation", logs: logs, sub: sub}, nil
}

// WatchAccountLiquidation is a free log subscription operation binding the contract event 0x4fa857769e6aa8d98c53ea164735f03821828e6672f227ea817ec1e74d61c343.
//
// Solidity: event AccountLiquidation(address indexed ownerAddress)
func (_Contract *ContractFilterer) WatchAccountLiquidation(opts *bind.WatchOpts, sink chan<- *ContractAccountLiquidation, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AccountLiquidation", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAccountLiquidation)
				if err := _Contract.contract.UnpackLog(event, "AccountLiquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountLiquidation is a log parse operation binding the contract event 0x4fa857769e6aa8d98c53ea164735f03821828e6672f227ea817ec1e74d61c343.
//
// Solidity: event AccountLiquidation(address indexed ownerAddress)
func (_Contract *ContractFilterer) ParseAccountLiquidation(log types.Log) (*ContractAccountLiquidation, error) {
	event := new(ContractAccountLiquidation)
	if err := _Contract.contract.UnpackLog(event, "AccountLiquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeclareOperatorFeePeriodUpdateIterator is returned from FilterDeclareOperatorFeePeriodUpdate and is used to iterate over the raw logs and unpacked data for DeclareOperatorFeePeriodUpdate events raised by the Contract contract.
type ContractDeclareOperatorFeePeriodUpdateIterator struct {
	Event *ContractDeclareOperatorFeePeriodUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeclareOperatorFeePeriodUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeclareOperatorFeePeriodUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeclareOperatorFeePeriodUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeclareOperatorFeePeriodUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeclareOperatorFeePeriodUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeclareOperatorFeePeriodUpdate represents a DeclareOperatorFeePeriodUpdate event raised by the Contract contract.
type ContractDeclareOperatorFeePeriodUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeclareOperatorFeePeriodUpdate is a free log retrieval operation binding the contract event 0xcaf8245d4ec588ab95e4db438be7dde364ed732720724028fefc713816d53711.
//
// Solidity: event DeclareOperatorFeePeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterDeclareOperatorFeePeriodUpdate(opts *bind.FilterOpts) (*ContractDeclareOperatorFeePeriodUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeclareOperatorFeePeriodUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractDeclareOperatorFeePeriodUpdateIterator{contract: _Contract.contract, event: "DeclareOperatorFeePeriodUpdate", logs: logs, sub: sub}, nil
}

// WatchDeclareOperatorFeePeriodUpdate is a free log subscription operation binding the contract event 0xcaf8245d4ec588ab95e4db438be7dde364ed732720724028fefc713816d53711.
//
// Solidity: event DeclareOperatorFeePeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchDeclareOperatorFeePeriodUpdate(opts *bind.WatchOpts, sink chan<- *ContractDeclareOperatorFeePeriodUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeclareOperatorFeePeriodUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeclareOperatorFeePeriodUpdate)
				if err := _Contract.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeclareOperatorFeePeriodUpdate is a log parse operation binding the contract event 0xcaf8245d4ec588ab95e4db438be7dde364ed732720724028fefc713816d53711.
//
// Solidity: event DeclareOperatorFeePeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseDeclareOperatorFeePeriodUpdate(log types.Log) (*ContractDeclareOperatorFeePeriodUpdate, error) {
	event := new(ContractDeclareOperatorFeePeriodUpdate)
	if err := _Contract.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeclaredOperatorFeeCancelationIterator is returned from FilterDeclaredOperatorFeeCancelation and is used to iterate over the raw logs and unpacked data for DeclaredOperatorFeeCancelation events raised by the Contract contract.
type ContractDeclaredOperatorFeeCancelationIterator struct {
	Event *ContractDeclaredOperatorFeeCancelation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeclaredOperatorFeeCancelationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeclaredOperatorFeeCancelation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeclaredOperatorFeeCancelation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeclaredOperatorFeeCancelationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeclaredOperatorFeeCancelationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeclaredOperatorFeeCancelation represents a DeclaredOperatorFeeCancelation event raised by the Contract contract.
type ContractDeclaredOperatorFeeCancelation struct {
	OwnerAddress common.Address
	OperatorId   uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeclaredOperatorFeeCancelation is a free log retrieval operation binding the contract event 0xb5c93858369b86b0bdf3a8c475316d42d2cd139306d28b3d5d4123b9b867455c.
//
// Solidity: event DeclaredOperatorFeeCancelation(address indexed ownerAddress, uint32 operatorId)
func (_Contract *ContractFilterer) FilterDeclaredOperatorFeeCancelation(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractDeclaredOperatorFeeCancelationIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeclaredOperatorFeeCancelation", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeclaredOperatorFeeCancelationIterator{contract: _Contract.contract, event: "DeclaredOperatorFeeCancelation", logs: logs, sub: sub}, nil
}

// WatchDeclaredOperatorFeeCancelation is a free log subscription operation binding the contract event 0xb5c93858369b86b0bdf3a8c475316d42d2cd139306d28b3d5d4123b9b867455c.
//
// Solidity: event DeclaredOperatorFeeCancelation(address indexed ownerAddress, uint32 operatorId)
func (_Contract *ContractFilterer) WatchDeclaredOperatorFeeCancelation(opts *bind.WatchOpts, sink chan<- *ContractDeclaredOperatorFeeCancelation, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeclaredOperatorFeeCancelation", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeclaredOperatorFeeCancelation)
				if err := _Contract.contract.UnpackLog(event, "DeclaredOperatorFeeCancelation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeclaredOperatorFeeCancelation is a log parse operation binding the contract event 0xb5c93858369b86b0bdf3a8c475316d42d2cd139306d28b3d5d4123b9b867455c.
//
// Solidity: event DeclaredOperatorFeeCancelation(address indexed ownerAddress, uint32 operatorId)
func (_Contract *ContractFilterer) ParseDeclaredOperatorFeeCancelation(log types.Log) (*ContractDeclaredOperatorFeeCancelation, error) {
	event := new(ContractDeclaredOperatorFeeCancelation)
	if err := _Contract.contract.UnpackLog(event, "DeclaredOperatorFeeCancelation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExecuteOperatorFeePeriodUpdateIterator is returned from FilterExecuteOperatorFeePeriodUpdate and is used to iterate over the raw logs and unpacked data for ExecuteOperatorFeePeriodUpdate events raised by the Contract contract.
type ContractExecuteOperatorFeePeriodUpdateIterator struct {
	Event *ContractExecuteOperatorFeePeriodUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractExecuteOperatorFeePeriodUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExecuteOperatorFeePeriodUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractExecuteOperatorFeePeriodUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractExecuteOperatorFeePeriodUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExecuteOperatorFeePeriodUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExecuteOperatorFeePeriodUpdate represents a ExecuteOperatorFeePeriodUpdate event raised by the Contract contract.
type ContractExecuteOperatorFeePeriodUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecuteOperatorFeePeriodUpdate is a free log retrieval operation binding the contract event 0x274c410ee0d5c156c7b7a46ff059b0df2394a5943e30bca4d54f01fb78a91c07.
//
// Solidity: event ExecuteOperatorFeePeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterExecuteOperatorFeePeriodUpdate(opts *bind.FilterOpts) (*ContractExecuteOperatorFeePeriodUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExecuteOperatorFeePeriodUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractExecuteOperatorFeePeriodUpdateIterator{contract: _Contract.contract, event: "ExecuteOperatorFeePeriodUpdate", logs: logs, sub: sub}, nil
}

// WatchExecuteOperatorFeePeriodUpdate is a free log subscription operation binding the contract event 0x274c410ee0d5c156c7b7a46ff059b0df2394a5943e30bca4d54f01fb78a91c07.
//
// Solidity: event ExecuteOperatorFeePeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchExecuteOperatorFeePeriodUpdate(opts *bind.WatchOpts, sink chan<- *ContractExecuteOperatorFeePeriodUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExecuteOperatorFeePeriodUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExecuteOperatorFeePeriodUpdate)
				if err := _Contract.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteOperatorFeePeriodUpdate is a log parse operation binding the contract event 0x274c410ee0d5c156c7b7a46ff059b0df2394a5943e30bca4d54f01fb78a91c07.
//
// Solidity: event ExecuteOperatorFeePeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseExecuteOperatorFeePeriodUpdate(log types.Log) (*ContractExecuteOperatorFeePeriodUpdate, error) {
	event := new(ContractExecuteOperatorFeePeriodUpdate)
	if err := _Contract.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFundsDepositIterator is returned from FilterFundsDeposit and is used to iterate over the raw logs and unpacked data for FundsDeposit events raised by the Contract contract.
type ContractFundsDepositIterator struct {
	Event *ContractFundsDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFundsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFundsDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFundsDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFundsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFundsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFundsDeposit represents a FundsDeposit event raised by the Contract contract.
type ContractFundsDeposit struct {
	Value         *big.Int
	OwnerAddress  common.Address
	SenderAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposit is a free log retrieval operation binding the contract event 0xb91f78378ad1ce7b5c12301896eac25f50d4603d0301406595e6fdfa6ffd6419.
//
// Solidity: event FundsDeposit(uint256 value, address indexed ownerAddress, address indexed senderAddress)
func (_Contract *ContractFilterer) FilterFundsDeposit(opts *bind.FilterOpts, ownerAddress []common.Address, senderAddress []common.Address) (*ContractFundsDepositIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}
	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FundsDeposit", ownerAddressRule, senderAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractFundsDepositIterator{contract: _Contract.contract, event: "FundsDeposit", logs: logs, sub: sub}, nil
}

// WatchFundsDeposit is a free log subscription operation binding the contract event 0xb91f78378ad1ce7b5c12301896eac25f50d4603d0301406595e6fdfa6ffd6419.
//
// Solidity: event FundsDeposit(uint256 value, address indexed ownerAddress, address indexed senderAddress)
func (_Contract *ContractFilterer) WatchFundsDeposit(opts *bind.WatchOpts, sink chan<- *ContractFundsDeposit, ownerAddress []common.Address, senderAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}
	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FundsDeposit", ownerAddressRule, senderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFundsDeposit)
				if err := _Contract.contract.UnpackLog(event, "FundsDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsDeposit is a log parse operation binding the contract event 0xb91f78378ad1ce7b5c12301896eac25f50d4603d0301406595e6fdfa6ffd6419.
//
// Solidity: event FundsDeposit(uint256 value, address indexed ownerAddress, address indexed senderAddress)
func (_Contract *ContractFilterer) ParseFundsDeposit(log types.Log) (*ContractFundsDeposit, error) {
	event := new(ContractFundsDeposit)
	if err := _Contract.contract.UnpackLog(event, "FundsDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFundsWithdrawalIterator is returned from FilterFundsWithdrawal and is used to iterate over the raw logs and unpacked data for FundsWithdrawal events raised by the Contract contract.
type ContractFundsWithdrawalIterator struct {
	Event *ContractFundsWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFundsWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFundsWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFundsWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFundsWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFundsWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFundsWithdrawal represents a FundsWithdrawal event raised by the Contract contract.
type ContractFundsWithdrawal struct {
	Value        *big.Int
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawal is a free log retrieval operation binding the contract event 0x233f5518707099d7be7f52de7c8094100dae65f0c896a762ebcc5900370d66bc.
//
// Solidity: event FundsWithdrawal(uint256 value, address indexed ownerAddress)
func (_Contract *ContractFilterer) FilterFundsWithdrawal(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractFundsWithdrawalIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FundsWithdrawal", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractFundsWithdrawalIterator{contract: _Contract.contract, event: "FundsWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawal is a free log subscription operation binding the contract event 0x233f5518707099d7be7f52de7c8094100dae65f0c896a762ebcc5900370d66bc.
//
// Solidity: event FundsWithdrawal(uint256 value, address indexed ownerAddress)
func (_Contract *ContractFilterer) WatchFundsWithdrawal(opts *bind.WatchOpts, sink chan<- *ContractFundsWithdrawal, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FundsWithdrawal", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFundsWithdrawal)
				if err := _Contract.contract.UnpackLog(event, "FundsWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsWithdrawal is a log parse operation binding the contract event 0x233f5518707099d7be7f52de7c8094100dae65f0c896a762ebcc5900370d66bc.
//
// Solidity: event FundsWithdrawal(uint256 value, address indexed ownerAddress)
func (_Contract *ContractFilterer) ParseFundsWithdrawal(log types.Log) (*ContractFundsWithdrawal, error) {
	event := new(ContractFundsWithdrawal)
	if err := _Contract.contract.UnpackLog(event, "FundsWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLiquidationThresholdPeriodUpdateIterator is returned from FilterLiquidationThresholdPeriodUpdate and is used to iterate over the raw logs and unpacked data for LiquidationThresholdPeriodUpdate events raised by the Contract contract.
type ContractLiquidationThresholdPeriodUpdateIterator struct {
	Event *ContractLiquidationThresholdPeriodUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractLiquidationThresholdPeriodUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLiquidationThresholdPeriodUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractLiquidationThresholdPeriodUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractLiquidationThresholdPeriodUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLiquidationThresholdPeriodUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLiquidationThresholdPeriodUpdate represents a LiquidationThresholdPeriodUpdate event raised by the Contract contract.
type ContractLiquidationThresholdPeriodUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidationThresholdPeriodUpdate is a free log retrieval operation binding the contract event 0x0ba69538f9785fc112824ff736c1e58ccf880282e09e53d7eef96598e282d3fa.
//
// Solidity: event LiquidationThresholdPeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterLiquidationThresholdPeriodUpdate(opts *bind.FilterOpts) (*ContractLiquidationThresholdPeriodUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LiquidationThresholdPeriodUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractLiquidationThresholdPeriodUpdateIterator{contract: _Contract.contract, event: "LiquidationThresholdPeriodUpdate", logs: logs, sub: sub}, nil
}

// WatchLiquidationThresholdPeriodUpdate is a free log subscription operation binding the contract event 0x0ba69538f9785fc112824ff736c1e58ccf880282e09e53d7eef96598e282d3fa.
//
// Solidity: event LiquidationThresholdPeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchLiquidationThresholdPeriodUpdate(opts *bind.WatchOpts, sink chan<- *ContractLiquidationThresholdPeriodUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LiquidationThresholdPeriodUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLiquidationThresholdPeriodUpdate)
				if err := _Contract.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidationThresholdPeriodUpdate is a log parse operation binding the contract event 0x0ba69538f9785fc112824ff736c1e58ccf880282e09e53d7eef96598e282d3fa.
//
// Solidity: event LiquidationThresholdPeriodUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseLiquidationThresholdPeriodUpdate(log types.Log) (*ContractLiquidationThresholdPeriodUpdate, error) {
	event := new(ContractLiquidationThresholdPeriodUpdate)
	if err := _Contract.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMinimumBlocksBeforeLiquidationUpdateIterator is returned from FilterMinimumBlocksBeforeLiquidationUpdate and is used to iterate over the raw logs and unpacked data for MinimumBlocksBeforeLiquidationUpdate events raised by the Contract contract.
type ContractMinimumBlocksBeforeLiquidationUpdateIterator struct {
	Event *ContractMinimumBlocksBeforeLiquidationUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMinimumBlocksBeforeLiquidationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMinimumBlocksBeforeLiquidationUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMinimumBlocksBeforeLiquidationUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMinimumBlocksBeforeLiquidationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMinimumBlocksBeforeLiquidationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMinimumBlocksBeforeLiquidationUpdate represents a MinimumBlocksBeforeLiquidationUpdate event raised by the Contract contract.
type ContractMinimumBlocksBeforeLiquidationUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinimumBlocksBeforeLiquidationUpdate is a free log retrieval operation binding the contract event 0xdc938b9b988f110956519739d5b987e32e22af85641fab70caa18d708d2ea9ff.
//
// Solidity: event MinimumBlocksBeforeLiquidationUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterMinimumBlocksBeforeLiquidationUpdate(opts *bind.FilterOpts) (*ContractMinimumBlocksBeforeLiquidationUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MinimumBlocksBeforeLiquidationUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractMinimumBlocksBeforeLiquidationUpdateIterator{contract: _Contract.contract, event: "MinimumBlocksBeforeLiquidationUpdate", logs: logs, sub: sub}, nil
}

// WatchMinimumBlocksBeforeLiquidationUpdate is a free log subscription operation binding the contract event 0xdc938b9b988f110956519739d5b987e32e22af85641fab70caa18d708d2ea9ff.
//
// Solidity: event MinimumBlocksBeforeLiquidationUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchMinimumBlocksBeforeLiquidationUpdate(opts *bind.WatchOpts, sink chan<- *ContractMinimumBlocksBeforeLiquidationUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MinimumBlocksBeforeLiquidationUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMinimumBlocksBeforeLiquidationUpdate)
				if err := _Contract.contract.UnpackLog(event, "MinimumBlocksBeforeLiquidationUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinimumBlocksBeforeLiquidationUpdate is a log parse operation binding the contract event 0xdc938b9b988f110956519739d5b987e32e22af85641fab70caa18d708d2ea9ff.
//
// Solidity: event MinimumBlocksBeforeLiquidationUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseMinimumBlocksBeforeLiquidationUpdate(log types.Log) (*ContractMinimumBlocksBeforeLiquidationUpdate, error) {
	event := new(ContractMinimumBlocksBeforeLiquidationUpdate)
	if err := _Contract.contract.UnpackLog(event, "MinimumBlocksBeforeLiquidationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNetworkFeeUpdateIterator is returned from FilterNetworkFeeUpdate and is used to iterate over the raw logs and unpacked data for NetworkFeeUpdate events raised by the Contract contract.
type ContractNetworkFeeUpdateIterator struct {
	Event *ContractNetworkFeeUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNetworkFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNetworkFeeUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNetworkFeeUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNetworkFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNetworkFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNetworkFeeUpdate represents a NetworkFeeUpdate event raised by the Contract contract.
type ContractNetworkFeeUpdate struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNetworkFeeUpdate is a free log retrieval operation binding the contract event 0xb12de2022a8c992455195d1cdfe49c528c96a325adc99e12b6eab59d56654ea2.
//
// Solidity: event NetworkFeeUpdate(uint256 oldFee, uint256 newFee)
func (_Contract *ContractFilterer) FilterNetworkFeeUpdate(opts *bind.FilterOpts) (*ContractNetworkFeeUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NetworkFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractNetworkFeeUpdateIterator{contract: _Contract.contract, event: "NetworkFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchNetworkFeeUpdate is a free log subscription operation binding the contract event 0xb12de2022a8c992455195d1cdfe49c528c96a325adc99e12b6eab59d56654ea2.
//
// Solidity: event NetworkFeeUpdate(uint256 oldFee, uint256 newFee)
func (_Contract *ContractFilterer) WatchNetworkFeeUpdate(opts *bind.WatchOpts, sink chan<- *ContractNetworkFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NetworkFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNetworkFeeUpdate)
				if err := _Contract.contract.UnpackLog(event, "NetworkFeeUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNetworkFeeUpdate is a log parse operation binding the contract event 0xb12de2022a8c992455195d1cdfe49c528c96a325adc99e12b6eab59d56654ea2.
//
// Solidity: event NetworkFeeUpdate(uint256 oldFee, uint256 newFee)
func (_Contract *ContractFilterer) ParseNetworkFeeUpdate(log types.Log) (*ContractNetworkFeeUpdate, error) {
	event := new(ContractNetworkFeeUpdate)
	if err := _Contract.contract.UnpackLog(event, "NetworkFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNetworkFeesWithdrawalIterator is returned from FilterNetworkFeesWithdrawal and is used to iterate over the raw logs and unpacked data for NetworkFeesWithdrawal events raised by the Contract contract.
type ContractNetworkFeesWithdrawalIterator struct {
	Event *ContractNetworkFeesWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNetworkFeesWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNetworkFeesWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNetworkFeesWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNetworkFeesWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNetworkFeesWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNetworkFeesWithdrawal represents a NetworkFeesWithdrawal event raised by the Contract contract.
type ContractNetworkFeesWithdrawal struct {
	Value     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNetworkFeesWithdrawal is a free log retrieval operation binding the contract event 0x5564de4b05294182a45c60836e9c0ba86d5c351cf1c0db79e4fe5dd04786de1a.
//
// Solidity: event NetworkFeesWithdrawal(uint256 value, address recipient)
func (_Contract *ContractFilterer) FilterNetworkFeesWithdrawal(opts *bind.FilterOpts) (*ContractNetworkFeesWithdrawalIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NetworkFeesWithdrawal")
	if err != nil {
		return nil, err
	}
	return &ContractNetworkFeesWithdrawalIterator{contract: _Contract.contract, event: "NetworkFeesWithdrawal", logs: logs, sub: sub}, nil
}

// WatchNetworkFeesWithdrawal is a free log subscription operation binding the contract event 0x5564de4b05294182a45c60836e9c0ba86d5c351cf1c0db79e4fe5dd04786de1a.
//
// Solidity: event NetworkFeesWithdrawal(uint256 value, address recipient)
func (_Contract *ContractFilterer) WatchNetworkFeesWithdrawal(opts *bind.WatchOpts, sink chan<- *ContractNetworkFeesWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NetworkFeesWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNetworkFeesWithdrawal)
				if err := _Contract.contract.UnpackLog(event, "NetworkFeesWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNetworkFeesWithdrawal is a log parse operation binding the contract event 0x5564de4b05294182a45c60836e9c0ba86d5c351cf1c0db79e4fe5dd04786de1a.
//
// Solidity: event NetworkFeesWithdrawal(uint256 value, address recipient)
func (_Contract *ContractFilterer) ParseNetworkFeesWithdrawal(log types.Log) (*ContractNetworkFeesWithdrawal, error) {
	event := new(ContractNetworkFeesWithdrawal)
	if err := _Contract.contract.UnpackLog(event, "NetworkFeesWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorFeeDeclarationIterator is returned from FilterOperatorFeeDeclaration and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclaration events raised by the Contract contract.
type ContractOperatorFeeDeclarationIterator struct {
	Event *ContractOperatorFeeDeclaration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorFeeDeclarationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorFeeDeclaration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorFeeDeclaration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorFeeDeclarationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorFeeDeclarationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorFeeDeclaration represents a OperatorFeeDeclaration event raised by the Contract contract.
type ContractOperatorFeeDeclaration struct {
	OwnerAddress common.Address
	OperatorId   uint32
	BlockNumber  *big.Int
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclaration is a free log retrieval operation binding the contract event 0x4892696cd76ab5f9fa4578747dee737eb9061eb170dab27ac24d78f47d3eec59.
//
// Solidity: event OperatorFeeDeclaration(address indexed ownerAddress, uint32 operatorId, uint256 blockNumber, uint256 fee)
func (_Contract *ContractFilterer) FilterOperatorFeeDeclaration(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractOperatorFeeDeclarationIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorFeeDeclaration", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorFeeDeclarationIterator{contract: _Contract.contract, event: "OperatorFeeDeclaration", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclaration is a free log subscription operation binding the contract event 0x4892696cd76ab5f9fa4578747dee737eb9061eb170dab27ac24d78f47d3eec59.
//
// Solidity: event OperatorFeeDeclaration(address indexed ownerAddress, uint32 operatorId, uint256 blockNumber, uint256 fee)
func (_Contract *ContractFilterer) WatchOperatorFeeDeclaration(opts *bind.WatchOpts, sink chan<- *ContractOperatorFeeDeclaration, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorFeeDeclaration", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorFeeDeclaration)
				if err := _Contract.contract.UnpackLog(event, "OperatorFeeDeclaration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeDeclaration is a log parse operation binding the contract event 0x4892696cd76ab5f9fa4578747dee737eb9061eb170dab27ac24d78f47d3eec59.
//
// Solidity: event OperatorFeeDeclaration(address indexed ownerAddress, uint32 operatorId, uint256 blockNumber, uint256 fee)
func (_Contract *ContractFilterer) ParseOperatorFeeDeclaration(log types.Log) (*ContractOperatorFeeDeclaration, error) {
	event := new(ContractOperatorFeeDeclaration)
	if err := _Contract.contract.UnpackLog(event, "OperatorFeeDeclaration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorFeeExecutionIterator is returned from FilterOperatorFeeExecution and is used to iterate over the raw logs and unpacked data for OperatorFeeExecution events raised by the Contract contract.
type ContractOperatorFeeExecutionIterator struct {
	Event *ContractOperatorFeeExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorFeeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorFeeExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorFeeExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorFeeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorFeeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorFeeExecution represents a OperatorFeeExecution event raised by the Contract contract.
type ContractOperatorFeeExecution struct {
	OwnerAddress common.Address
	OperatorId   uint32
	BlockNumber  *big.Int
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeExecution is a free log retrieval operation binding the contract event 0x5e68243d185e61c61d839a6eb781a865dd007c0004b0c1f5f544e4291e7ee96d.
//
// Solidity: event OperatorFeeExecution(address indexed ownerAddress, uint32 operatorId, uint256 blockNumber, uint256 fee)
func (_Contract *ContractFilterer) FilterOperatorFeeExecution(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractOperatorFeeExecutionIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorFeeExecution", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorFeeExecutionIterator{contract: _Contract.contract, event: "OperatorFeeExecution", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeExecution is a free log subscription operation binding the contract event 0x5e68243d185e61c61d839a6eb781a865dd007c0004b0c1f5f544e4291e7ee96d.
//
// Solidity: event OperatorFeeExecution(address indexed ownerAddress, uint32 operatorId, uint256 blockNumber, uint256 fee)
func (_Contract *ContractFilterer) WatchOperatorFeeExecution(opts *bind.WatchOpts, sink chan<- *ContractOperatorFeeExecution, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorFeeExecution", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorFeeExecution)
				if err := _Contract.contract.UnpackLog(event, "OperatorFeeExecution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeExecution is a log parse operation binding the contract event 0x5e68243d185e61c61d839a6eb781a865dd007c0004b0c1f5f544e4291e7ee96d.
//
// Solidity: event OperatorFeeExecution(address indexed ownerAddress, uint32 operatorId, uint256 blockNumber, uint256 fee)
func (_Contract *ContractFilterer) ParseOperatorFeeExecution(log types.Log) (*ContractOperatorFeeExecution, error) {
	event := new(ContractOperatorFeeExecution)
	if err := _Contract.contract.UnpackLog(event, "OperatorFeeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorFeeIncreaseLimitUpdateIterator is returned from FilterOperatorFeeIncreaseLimitUpdate and is used to iterate over the raw logs and unpacked data for OperatorFeeIncreaseLimitUpdate events raised by the Contract contract.
type ContractOperatorFeeIncreaseLimitUpdateIterator struct {
	Event *ContractOperatorFeeIncreaseLimitUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorFeeIncreaseLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorFeeIncreaseLimitUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorFeeIncreaseLimitUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorFeeIncreaseLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorFeeIncreaseLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorFeeIncreaseLimitUpdate represents a OperatorFeeIncreaseLimitUpdate event raised by the Contract contract.
type ContractOperatorFeeIncreaseLimitUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeIncreaseLimitUpdate is a free log retrieval operation binding the contract event 0x65703aa387fec070bf1ec4b824c8faee64d3ad3e45b03c119727391b06d168a7.
//
// Solidity: event OperatorFeeIncreaseLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterOperatorFeeIncreaseLimitUpdate(opts *bind.FilterOpts) (*ContractOperatorFeeIncreaseLimitUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorFeeIncreaseLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractOperatorFeeIncreaseLimitUpdateIterator{contract: _Contract.contract, event: "OperatorFeeIncreaseLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeIncreaseLimitUpdate is a free log subscription operation binding the contract event 0x65703aa387fec070bf1ec4b824c8faee64d3ad3e45b03c119727391b06d168a7.
//
// Solidity: event OperatorFeeIncreaseLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchOperatorFeeIncreaseLimitUpdate(opts *bind.WatchOpts, sink chan<- *ContractOperatorFeeIncreaseLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorFeeIncreaseLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorFeeIncreaseLimitUpdate)
				if err := _Contract.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeIncreaseLimitUpdate is a log parse operation binding the contract event 0x65703aa387fec070bf1ec4b824c8faee64d3ad3e45b03c119727391b06d168a7.
//
// Solidity: event OperatorFeeIncreaseLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseOperatorFeeIncreaseLimitUpdate(log types.Log) (*ContractOperatorFeeIncreaseLimitUpdate, error) {
	event := new(ContractOperatorFeeIncreaseLimitUpdate)
	if err := _Contract.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorMaxFeeIncreaseUpdateIterator is returned from FilterOperatorMaxFeeIncreaseUpdate and is used to iterate over the raw logs and unpacked data for OperatorMaxFeeIncreaseUpdate events raised by the Contract contract.
type ContractOperatorMaxFeeIncreaseUpdateIterator struct {
	Event *ContractOperatorMaxFeeIncreaseUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorMaxFeeIncreaseUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorMaxFeeIncreaseUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorMaxFeeIncreaseUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorMaxFeeIncreaseUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorMaxFeeIncreaseUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorMaxFeeIncreaseUpdate represents a OperatorMaxFeeIncreaseUpdate event raised by the Contract contract.
type ContractOperatorMaxFeeIncreaseUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOperatorMaxFeeIncreaseUpdate is a free log retrieval operation binding the contract event 0xe36a3666a349caa56ad2348c16813da281ed07220379997887a71b790f0f0a47.
//
// Solidity: event OperatorMaxFeeIncreaseUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterOperatorMaxFeeIncreaseUpdate(opts *bind.FilterOpts) (*ContractOperatorMaxFeeIncreaseUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorMaxFeeIncreaseUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractOperatorMaxFeeIncreaseUpdateIterator{contract: _Contract.contract, event: "OperatorMaxFeeIncreaseUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorMaxFeeIncreaseUpdate is a free log subscription operation binding the contract event 0xe36a3666a349caa56ad2348c16813da281ed07220379997887a71b790f0f0a47.
//
// Solidity: event OperatorMaxFeeIncreaseUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchOperatorMaxFeeIncreaseUpdate(opts *bind.WatchOpts, sink chan<- *ContractOperatorMaxFeeIncreaseUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorMaxFeeIncreaseUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorMaxFeeIncreaseUpdate)
				if err := _Contract.contract.UnpackLog(event, "OperatorMaxFeeIncreaseUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorMaxFeeIncreaseUpdate is a log parse operation binding the contract event 0xe36a3666a349caa56ad2348c16813da281ed07220379997887a71b790f0f0a47.
//
// Solidity: event OperatorMaxFeeIncreaseUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseOperatorMaxFeeIncreaseUpdate(log types.Log) (*ContractOperatorMaxFeeIncreaseUpdate, error) {
	event := new(ContractOperatorMaxFeeIncreaseUpdate)
	if err := _Contract.contract.UnpackLog(event, "OperatorMaxFeeIncreaseUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorRegistrationIterator is returned from FilterOperatorRegistration and is used to iterate over the raw logs and unpacked data for OperatorRegistration events raised by the Contract contract.
type ContractOperatorRegistrationIterator struct {
	Event *ContractOperatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorRegistration represents a OperatorRegistration event raised by the Contract contract.
type ContractOperatorRegistration struct {
	Id           uint32
	Name         string
	OwnerAddress common.Address
	PublicKey    []byte
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistration is a free log retrieval operation binding the contract event 0x26a77904793977b23eb8b2d412c486276510e0dc1966a4a2936d4bea0ff86e9d.
//
// Solidity: event OperatorRegistration(uint32 indexed id, string name, address indexed ownerAddress, bytes publicKey, uint256 fee)
func (_Contract *ContractFilterer) FilterOperatorRegistration(opts *bind.FilterOpts, id []uint32, ownerAddress []common.Address) (*ContractOperatorRegistrationIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorRegistration", idRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorRegistrationIterator{contract: _Contract.contract, event: "OperatorRegistration", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistration is a free log subscription operation binding the contract event 0x26a77904793977b23eb8b2d412c486276510e0dc1966a4a2936d4bea0ff86e9d.
//
// Solidity: event OperatorRegistration(uint32 indexed id, string name, address indexed ownerAddress, bytes publicKey, uint256 fee)
func (_Contract *ContractFilterer) WatchOperatorRegistration(opts *bind.WatchOpts, sink chan<- *ContractOperatorRegistration, id []uint32, ownerAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorRegistration", idRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorRegistration)
				if err := _Contract.contract.UnpackLog(event, "OperatorRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistration is a log parse operation binding the contract event 0x26a77904793977b23eb8b2d412c486276510e0dc1966a4a2936d4bea0ff86e9d.
//
// Solidity: event OperatorRegistration(uint32 indexed id, string name, address indexed ownerAddress, bytes publicKey, uint256 fee)
func (_Contract *ContractFilterer) ParseOperatorRegistration(log types.Log) (*ContractOperatorRegistration, error) {
	event := new(ContractOperatorRegistration)
	if err := _Contract.contract.UnpackLog(event, "OperatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorRemovalIterator is returned from FilterOperatorRemoval and is used to iterate over the raw logs and unpacked data for OperatorRemoval events raised by the Contract contract.
type ContractOperatorRemovalIterator struct {
	Event *ContractOperatorRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorRemoval represents a OperatorRemoval event raised by the Contract contract.
type ContractOperatorRemoval struct {
	OperatorId   uint32
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoval is a free log retrieval operation binding the contract event 0x10b90e3d042178ee9b3e99a849224b4bf4145b9855274073c0c6bca9c5113b7b.
//
// Solidity: event OperatorRemoval(uint32 operatorId, address indexed ownerAddress)
func (_Contract *ContractFilterer) FilterOperatorRemoval(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractOperatorRemovalIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorRemoval", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorRemovalIterator{contract: _Contract.contract, event: "OperatorRemoval", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoval is a free log subscription operation binding the contract event 0x10b90e3d042178ee9b3e99a849224b4bf4145b9855274073c0c6bca9c5113b7b.
//
// Solidity: event OperatorRemoval(uint32 operatorId, address indexed ownerAddress)
func (_Contract *ContractFilterer) WatchOperatorRemoval(opts *bind.WatchOpts, sink chan<- *ContractOperatorRemoval, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorRemoval", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorRemoval)
				if err := _Contract.contract.UnpackLog(event, "OperatorRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRemoval is a log parse operation binding the contract event 0x10b90e3d042178ee9b3e99a849224b4bf4145b9855274073c0c6bca9c5113b7b.
//
// Solidity: event OperatorRemoval(uint32 operatorId, address indexed ownerAddress)
func (_Contract *ContractFilterer) ParseOperatorRemoval(log types.Log) (*ContractOperatorRemoval, error) {
	event := new(ContractOperatorRemoval)
	if err := _Contract.contract.UnpackLog(event, "OperatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOperatorScoreUpdateIterator is returned from FilterOperatorScoreUpdate and is used to iterate over the raw logs and unpacked data for OperatorScoreUpdate events raised by the Contract contract.
type ContractOperatorScoreUpdateIterator struct {
	Event *ContractOperatorScoreUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOperatorScoreUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorScoreUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOperatorScoreUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOperatorScoreUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOperatorScoreUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOperatorScoreUpdate represents a OperatorScoreUpdate event raised by the Contract contract.
type ContractOperatorScoreUpdate struct {
	OperatorId   uint32
	OwnerAddress common.Address
	BlockNumber  *big.Int
	Score        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorScoreUpdate is a free log retrieval operation binding the contract event 0xc07e279e1b82e9c5a89dc92f0d0fb3ec87c142b3cc1a27411ab82156beef8c36.
//
// Solidity: event OperatorScoreUpdate(uint32 operatorId, address indexed ownerAddress, uint256 blockNumber, uint256 score)
func (_Contract *ContractFilterer) FilterOperatorScoreUpdate(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractOperatorScoreUpdateIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorScoreUpdate", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorScoreUpdateIterator{contract: _Contract.contract, event: "OperatorScoreUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorScoreUpdate is a free log subscription operation binding the contract event 0xc07e279e1b82e9c5a89dc92f0d0fb3ec87c142b3cc1a27411ab82156beef8c36.
//
// Solidity: event OperatorScoreUpdate(uint32 operatorId, address indexed ownerAddress, uint256 blockNumber, uint256 score)
func (_Contract *ContractFilterer) WatchOperatorScoreUpdate(opts *bind.WatchOpts, sink chan<- *ContractOperatorScoreUpdate, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorScoreUpdate", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOperatorScoreUpdate)
				if err := _Contract.contract.UnpackLog(event, "OperatorScoreUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorScoreUpdate is a log parse operation binding the contract event 0xc07e279e1b82e9c5a89dc92f0d0fb3ec87c142b3cc1a27411ab82156beef8c36.
//
// Solidity: event OperatorScoreUpdate(uint32 operatorId, address indexed ownerAddress, uint256 blockNumber, uint256 score)
func (_Contract *ContractFilterer) ParseOperatorScoreUpdate(log types.Log) (*ContractOperatorScoreUpdate, error) {
	event := new(ContractOperatorScoreUpdate)
	if err := _Contract.contract.UnpackLog(event, "OperatorScoreUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegisteredOperatorsPerAccountLimitUpdateIterator is returned from FilterRegisteredOperatorsPerAccountLimitUpdate and is used to iterate over the raw logs and unpacked data for RegisteredOperatorsPerAccountLimitUpdate events raised by the Contract contract.
type ContractRegisteredOperatorsPerAccountLimitUpdateIterator struct {
	Event *ContractRegisteredOperatorsPerAccountLimitUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRegisteredOperatorsPerAccountLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegisteredOperatorsPerAccountLimitUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRegisteredOperatorsPerAccountLimitUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRegisteredOperatorsPerAccountLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegisteredOperatorsPerAccountLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegisteredOperatorsPerAccountLimitUpdate represents a RegisteredOperatorsPerAccountLimitUpdate event raised by the Contract contract.
type ContractRegisteredOperatorsPerAccountLimitUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegisteredOperatorsPerAccountLimitUpdate is a free log retrieval operation binding the contract event 0xf54ba6bfb6e718f7df9254b410fc215ac507e5a054c844e35856b943f71356b6.
//
// Solidity: event RegisteredOperatorsPerAccountLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterRegisteredOperatorsPerAccountLimitUpdate(opts *bind.FilterOpts) (*ContractRegisteredOperatorsPerAccountLimitUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RegisteredOperatorsPerAccountLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractRegisteredOperatorsPerAccountLimitUpdateIterator{contract: _Contract.contract, event: "RegisteredOperatorsPerAccountLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchRegisteredOperatorsPerAccountLimitUpdate is a free log subscription operation binding the contract event 0xf54ba6bfb6e718f7df9254b410fc215ac507e5a054c844e35856b943f71356b6.
//
// Solidity: event RegisteredOperatorsPerAccountLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchRegisteredOperatorsPerAccountLimitUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegisteredOperatorsPerAccountLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RegisteredOperatorsPerAccountLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegisteredOperatorsPerAccountLimitUpdate)
				if err := _Contract.contract.UnpackLog(event, "RegisteredOperatorsPerAccountLimitUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisteredOperatorsPerAccountLimitUpdate is a log parse operation binding the contract event 0xf54ba6bfb6e718f7df9254b410fc215ac507e5a054c844e35856b943f71356b6.
//
// Solidity: event RegisteredOperatorsPerAccountLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseRegisteredOperatorsPerAccountLimitUpdate(log types.Log) (*ContractRegisteredOperatorsPerAccountLimitUpdate, error) {
	event := new(ContractRegisteredOperatorsPerAccountLimitUpdate)
	if err := _Contract.contract.UnpackLog(event, "RegisteredOperatorsPerAccountLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorRegistrationIterator is returned from FilterValidatorRegistration and is used to iterate over the raw logs and unpacked data for ValidatorRegistration events raised by the Contract contract.
type ContractValidatorRegistrationIterator struct {
	Event *ContractValidatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorRegistration represents a ValidatorRegistration event raised by the Contract contract.
type ContractValidatorRegistration struct {
	OwnerAddress     common.Address
	PublicKey        []byte
	OperatorIds      []uint32
	SharesPublicKeys [][]byte
	EncryptedKeys    [][]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistration is a free log retrieval operation binding the contract event 0x888b4bb563730efc1c420fb22b503c3551134948a3a3dce4ffab6380e9ce5025.
//
// Solidity: event ValidatorRegistration(address indexed ownerAddress, bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] encryptedKeys)
func (_Contract *ContractFilterer) FilterValidatorRegistration(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractValidatorRegistrationIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorRegistration", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorRegistrationIterator{contract: _Contract.contract, event: "ValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistration is a free log subscription operation binding the contract event 0x888b4bb563730efc1c420fb22b503c3551134948a3a3dce4ffab6380e9ce5025.
//
// Solidity: event ValidatorRegistration(address indexed ownerAddress, bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] encryptedKeys)
func (_Contract *ContractFilterer) WatchValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ContractValidatorRegistration, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorRegistration", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorRegistration)
				if err := _Contract.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistration is a log parse operation binding the contract event 0x888b4bb563730efc1c420fb22b503c3551134948a3a3dce4ffab6380e9ce5025.
//
// Solidity: event ValidatorRegistration(address indexed ownerAddress, bytes publicKey, uint32[] operatorIds, bytes[] sharesPublicKeys, bytes[] encryptedKeys)
func (_Contract *ContractFilterer) ParseValidatorRegistration(log types.Log) (*ContractValidatorRegistration, error) {
	event := new(ContractValidatorRegistration)
	if err := _Contract.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorRemovalIterator is returned from FilterValidatorRemoval and is used to iterate over the raw logs and unpacked data for ValidatorRemoval events raised by the Contract contract.
type ContractValidatorRemovalIterator struct {
	Event *ContractValidatorRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorRemoval represents a ValidatorRemoval event raised by the Contract contract.
type ContractValidatorRemoval struct {
	OwnerAddress common.Address
	PublicKey    []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoval is a free log retrieval operation binding the contract event 0x671ada3835502b9498e4a3116c344293ec3a4ef43f90bb42283d7d66a3f772b2.
//
// Solidity: event ValidatorRemoval(address indexed ownerAddress, bytes publicKey)
func (_Contract *ContractFilterer) FilterValidatorRemoval(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractValidatorRemovalIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorRemoval", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorRemovalIterator{contract: _Contract.contract, event: "ValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoval is a free log subscription operation binding the contract event 0x671ada3835502b9498e4a3116c344293ec3a4ef43f90bb42283d7d66a3f772b2.
//
// Solidity: event ValidatorRemoval(address indexed ownerAddress, bytes publicKey)
func (_Contract *ContractFilterer) WatchValidatorRemoval(opts *bind.WatchOpts, sink chan<- *ContractValidatorRemoval, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorRemoval", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorRemoval)
				if err := _Contract.contract.UnpackLog(event, "ValidatorRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemoval is a log parse operation binding the contract event 0x671ada3835502b9498e4a3116c344293ec3a4ef43f90bb42283d7d66a3f772b2.
//
// Solidity: event ValidatorRemoval(address indexed ownerAddress, bytes publicKey)
func (_Contract *ContractFilterer) ParseValidatorRemoval(log types.Log) (*ContractValidatorRemoval, error) {
	event := new(ContractValidatorRemoval)
	if err := _Contract.contract.UnpackLog(event, "ValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorsPerOperatorLimitUpdateIterator is returned from FilterValidatorsPerOperatorLimitUpdate and is used to iterate over the raw logs and unpacked data for ValidatorsPerOperatorLimitUpdate events raised by the Contract contract.
type ContractValidatorsPerOperatorLimitUpdateIterator struct {
	Event *ContractValidatorsPerOperatorLimitUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorsPerOperatorLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorsPerOperatorLimitUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorsPerOperatorLimitUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorsPerOperatorLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorsPerOperatorLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorsPerOperatorLimitUpdate represents a ValidatorsPerOperatorLimitUpdate event raised by the Contract contract.
type ContractValidatorsPerOperatorLimitUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorsPerOperatorLimitUpdate is a free log retrieval operation binding the contract event 0x8f74111035663610cd2e52bd3b3c4af95e44815380fd78eb37de9775a3824a3b.
//
// Solidity: event ValidatorsPerOperatorLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) FilterValidatorsPerOperatorLimitUpdate(opts *bind.FilterOpts) (*ContractValidatorsPerOperatorLimitUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorsPerOperatorLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractValidatorsPerOperatorLimitUpdateIterator{contract: _Contract.contract, event: "ValidatorsPerOperatorLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorsPerOperatorLimitUpdate is a free log subscription operation binding the contract event 0x8f74111035663610cd2e52bd3b3c4af95e44815380fd78eb37de9775a3824a3b.
//
// Solidity: event ValidatorsPerOperatorLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) WatchValidatorsPerOperatorLimitUpdate(opts *bind.WatchOpts, sink chan<- *ContractValidatorsPerOperatorLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorsPerOperatorLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorsPerOperatorLimitUpdate)
				if err := _Contract.contract.UnpackLog(event, "ValidatorsPerOperatorLimitUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsPerOperatorLimitUpdate is a log parse operation binding the contract event 0x8f74111035663610cd2e52bd3b3c4af95e44815380fd78eb37de9775a3824a3b.
//
// Solidity: event ValidatorsPerOperatorLimitUpdate(uint256 value)
func (_Contract *ContractFilterer) ParseValidatorsPerOperatorLimitUpdate(log types.Log) (*ContractValidatorsPerOperatorLimitUpdate, error) {
	event := new(ContractValidatorsPerOperatorLimitUpdate)
	if err := _Contract.contract.UnpackLog(event, "ValidatorsPerOperatorLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
