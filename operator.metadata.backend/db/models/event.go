package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OperatorAddedEvent struct {
	OperatorId uint64         // indexed
	Owner      common.Address // indexed
	PublicKey  []byte
	Fee        *big.Int
}
