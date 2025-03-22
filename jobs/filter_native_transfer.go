package jobs

import (
	"context"
	"eth-watcher/config"
	"eth-watcher/global"
	"eth-watcher/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func isInAddressArray(target ethCommon.Address, arr []ethCommon.Address) bool {
	for _, addr := range arr {
		if addr.Cmp(target) == 0 {
			return true
		}
	}
	return false
}

func filterNativeTransfer(client *ethclient.Client, fromBlock uint64, toBlock uint64) (filteredLogs []types.FilterParsedLog, err error) {

	if fromBlock > toBlock {
		// Just nothing
		return nil, nil
	}

	// Inspect every block
	for blockNo := fromBlock; blockNo <= toBlock; blockNo++ {
		global.Logger.Debugf("Inspecting block #%d ...", blockNo)
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
			// Check transaction's sender and receiver

			signer, err := ethTypes.Sender(ethTypes.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil {
				global.Logger.Errorf("Failed to extract signer from transaction %s with error: %v", tx.Hash(), err)
				break
			}

			receiver := *tx.To()

			var ia *ethCommon.Address = nil
			if isInAddressArray(signer, config.Config.SendersAddress) {
				ia = &signer
			} else if isInAddressArray(receiver, config.Config.ReceiversAddress) {
				ia = &receiver
			}

			if ia != nil {
				// Found
				filteredLogs = append(filteredLogs, types.FilterParsedLog{
					IAddress:    *ia,
					BlockNumber: block.Number(),
					Sender:      signer,
					Receiver:    receiver,
					Amount:      tx.Value(),
					Contract:    nil,
					TxHash:      tx.Hash(),
					TimeStamp:   time.Unix(int64(block.Time()), 0),
				})
			}
		}
	}

	return filteredLogs, nil

}
