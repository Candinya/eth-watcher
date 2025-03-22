package utils

import (
	"context"
	"eth-watcher/chain"
	"eth-watcher/global"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FilterCallback(chainCfg *types.ChainConfig, isNative bool, client *ethclient.Client, parsedLog types.FilterParsedLog) {
	if isNative {
		// Parse amount
		amountParsed := ParseAmount(parsedLog.Amount, nil)

		balance, err := client.BalanceAt(context.Background(), parsedLog.IAddress, parsedLog.BlockNumber)
		if err != nil {
			global.Logger.Errorf("Failed to get balance for accout %s with error: %v", parsedLog.IAddress.Hex(), err)
			return
		}

		balanceParsed := ParseAmount(balance, nil)

		// Call webhook
		WebhookCallback(chainCfg, parsedLog.IAddress.Hex(), parsedLog.Sender.Hex(), parsedLog.Receiver.Hex(), true, nil, nil, amountParsed, balanceParsed, parsedLog.TxHash.Hex(), parsedLog.TimeStamp)
	} else {
		// Get contract meta
		contractMeta, err := GetERC20ContractMeta(chainCfg.ID, parsedLog.Contract, client)
		if err != nil {
			global.Logger.Errorf("Failed to get meta of contract %s with error: %v", parsedLog.Contract.Hex(), err)
			return
		}

		// Parse amount
		amountParsed := ParseAmount(parsedLog.Amount, &contractMeta.Decimals)

		// Get current balance
		balance, err := chain.GetERC20Balance(parsedLog.Contract, &parsedLog.IAddress, parsedLog.BlockNumber, client)
		if err != nil {
			global.Logger.Errorf("Failed to get balance for accout %s of contract %s with error: %v", parsedLog.IAddress.Hex(), parsedLog.Contract.Hex(), err)
			return
		}

		balanceParsed := ParseAmount(balance, &contractMeta.Decimals)

		// Call webhook
		WebhookCallback(chainCfg, parsedLog.IAddress.Hex(), parsedLog.Sender.Hex(), parsedLog.Receiver.Hex(), false, P(parsedLog.Contract.Hex()), contractMeta, amountParsed, balanceParsed, parsedLog.TxHash.Hex(), parsedLog.TimeStamp)
	}
}
