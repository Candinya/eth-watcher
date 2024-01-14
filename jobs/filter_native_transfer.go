package jobs

import (
	"context"
	"eth-watcher/config"
	"eth-watcher/global"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func filterNativeTransfer(client *ethclient.Client, fromBlock uint64, toBlock uint64) (filteredLogs []types.FilterParsedLog, err error) {

	if fromBlock > toBlock {
		// Just nothing
		return nil, nil
	}

	// Inspect every block
	for blockNo := fromBlock; blockNo <= toBlock; blockNo++ {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNo)))
		if err != nil {
			global.Logger.Errorf("Failed to inspect block #%d with error: %v", blockNo, err)
			continue
		}
		// Check every transaction
		for _, tx := range block.Transactions() {
			if tx.To() == nil || // Is contract creation
				tx.Value().Cmp(big.NewInt(0)) == 0 { // Nothing transferred
				continue
			}
			// Check if transaction recipient is in receivers
			toAddressHash := common.BytesToHash(tx.To().Bytes())
			for _, receiver := range config.Config.ReceiversHash {
				if receiver.Cmp(toAddressHash) == 0 {
					// Found
					sender, err := ethTypes.Sender(ethTypes.LatestSignerForChainID(tx.ChainId()), tx)
					if err != nil {
						global.Logger.Errorf("Failed to extract signer from transaction %s with error: %v", tx.Hash(), err)
						break
					}

					filteredLogs = append(filteredLogs, types.FilterParsedLog{
						Sender:    sender,
						Receiver:  *tx.To(),
						Amount:    tx.Value(),
						Contract:  nil,
						TxHash:    tx.Hash(),
						TimeStamp: time.Unix(int64(block.Time()), 0),
					})

					break
				}
			}
		}
	}

	return filteredLogs, nil

}
