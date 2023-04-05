package utils

import (
	"eth-watcher/global"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FilterCallback(chain *types.ChainConfig, isNative bool, client *ethclient.Client, parsedLog types.FilterParsedLog) {
	// Get contract meta
	contractMeta, err := GetERC20ContractMeta(chain.ID, parsedLog.Contract, client)
	if err != nil {
		global.Logger.Errorf("Failed to get meta of contract %s with error: %v", parsedLog.Contract.Hex(), err)
		return
	}

	// Parse amount
	amountParsed := ParseAmount(parsedLog.Amount, contractMeta)

	// Call webhook
	WebhookCallback(chain, parsedLog.Sender.Hex(), parsedLog.Receiver.Hex(), isNative, contractMeta, amountParsed, parsedLog.TxHash.Hex())
}
