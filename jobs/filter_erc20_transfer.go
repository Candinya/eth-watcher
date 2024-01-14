package jobs

import (
	"context"
	"eth-watcher/config"
	"eth-watcher/consts"
	"eth-watcher/global"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func filterERC20Transfer(client *ethclient.Client, fromBlock uint64, toBlock uint64, contractWhitelist []ethCommon.Address) (filteredLogs []types.FilterParsedLog, err error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: contractWhitelist,
		Topics: [][]ethCommon.Hash{
			{consts.ERC20_TRANSFER_TOPIC_0},
			{},                          // From any
			config.Config.ReceiversHash, // To receivers
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		global.Logger.Errorf("Failed to filter log for block ( %d - %d ) with error: %v", fromBlock, toBlock, err)
		return nil, err
	}

	blockTs := make(map[uint64]time.Time)

	// Get logs (refer to ERC20 Token Transfer Topic definitions)
	zeroHash := ethCommon.BigToHash(big.NewInt(0))
	for _, log := range logs {
		if log.Topics[1].Cmp(zeroHash) != 0 {
			// Not mint (mint is from null (genesis) address)
			// Set timestamp
			ts, ok := blockTs[log.BlockNumber]
			if !ok {
				bInfo, err := client.BlockByNumber(context.Background(), big.NewInt(int64(log.BlockNumber)))
				if err != nil {
					global.Logger.Errorf("Failed to get info for block %d with error: %v", log.BlockNumber, err)
					ts = time.Unix(0, 0)
				} else {
					ts = time.Unix(int64(bInfo.Time()), 0)
					blockTs[log.BlockNumber] = ts
				}
			}
			filteredLogs = append(filteredLogs, types.FilterParsedLog{
				Sender:    ethCommon.BytesToAddress(log.Topics[1].Bytes()),
				Receiver:  ethCommon.BytesToAddress(log.Topics[2].Bytes()),
				Amount:    new(big.Int).SetBytes(log.Data),
				Contract:  &log.Address,
				TxHash:    log.TxHash,
				TimeStamp: ts,
			})
		}
	}

	return filteredLogs, nil
}
