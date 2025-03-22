package chain

import (
	"eth-watcher/chain/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetERC20Balance(contractAddress *ethCommon.Address, account *ethCommon.Address, blockNumber *big.Int, client *ethclient.Client) (*big.Int, error) {
	// Initialize contract
	instance, err := erc20.NewErc20(*contractAddress, client)
	if err != nil {
		return nil, err
	}

	return instance.BalanceOf(&bind.CallOpts{
		BlockNumber: blockNumber,
	}, *account)
}
