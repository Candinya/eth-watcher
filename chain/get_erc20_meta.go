package chain

import (
	"eth-watcher/chain/erc20"
	"eth-watcher/consts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetERC20Meta(contractAddress *ethCommon.Address, client *ethclient.Client) (name string, symbol string, decimals int64, err error) {
	// Initialize contract
	instance, err := erc20.NewErc20(*contractAddress, client)
	if err != nil {
		return "", "", consts.NATIVE_DECIMALS, err
	}

	// Call contract
	name, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		return "", "", consts.NATIVE_DECIMALS, err
	}

	symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", consts.NATIVE_DECIMALS, err
	}

	decimalsBig, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return "", "", consts.NATIVE_DECIMALS, err
	}
	decimals = int64(decimalsBig)

	return name, symbol, decimals, nil

}
