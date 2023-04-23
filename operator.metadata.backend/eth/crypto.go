package eth

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

// VerifyTypeDataSignV4 eip 712
func VerifyTypeDataSignV4(data string, signature string, address common.Address) (bool, error) {
	verified, _, err := VerifyTypeDataSignV4Plus(data, signature, address)
	return verified, err
}

func VerifyTypeDataSignV4Plus(data string, signature string, address common.Address) (bool, *apitypes.TypedData, error) {
	//	var data = `
	//   {"domain":{"name":"Ether Mail","chainId":5},"message":{"contents":"Hello, Bob!"},"primaryType":"Mail","types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"chainId","type":"uint256"}],"Mail":[{"name":"contents","type":"string"}]}}
	//`
	var d apitypes.TypedData
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		fmt.Println("=====VerifyTypeDataSignV4Plus Unmarshal====", err)
		return false, &d, err
	}
	verified, err := sigverify.VerifyTypedDataHexSignatureEx(address, d, signature)

	if err != nil {
		return false, &d, err
	}
	return verified, &d, err
}
