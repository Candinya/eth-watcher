package types

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type FilterParsedLog struct {
	IAddress    ethCommon.Address // Interest address
	BlockNumber *big.Int
	Sender      ethCommon.Address
	Receiver    ethCommon.Address
	Amount      *big.Int
	Contract    *ethCommon.Address
	TxHash      ethCommon.Hash
	TimeStamp   time.Time
}
