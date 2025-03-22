package utils

import (
	"eth-watcher/consts"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

func ParseAmount(rawAmount *big.Int, amountDecimals *int64) float64 {
	var decimals int64 = consts.NATIVE_DECIMALS
	if amountDecimals != nil {
		decimals = *amountDecimals
	}

	result, _ := new(big.Float).Quo(
		new(big.Float).SetInt(rawAmount),
		new(big.Float).SetInt(
			math.BigPow(10, decimals),
		),
	).Float64()

	return result // Ignore accuracy
}
