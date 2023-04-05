package utils

import (
	"eth-watcher/consts"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

func ParseAmount(rawAmount *big.Int, contractMeta *types.ContractMeta) float64 {
	var decimals int64 = consts.NATIVE_DECIMALS
	if contractMeta != nil {
		decimals = contractMeta.Decimals
	}

	result, _ := new(big.Float).Quo(
		new(big.Float).SetInt(rawAmount),
		new(big.Float).SetInt(
			math.Exp(big.NewInt(10), big.NewInt(decimals)),
		),
	).Float64()

	return result // Ignore accuracy
}
