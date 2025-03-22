package jobs

import (
	"context"
	"eth-watcher/config"
	"eth-watcher/consts"
	"eth-watcher/global"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type logWithIA struct {
	IA ethCommon.Address
	L  ethTypes.Log
}

func filterERC20Transfer(client *ethclient.Client, fromBlock uint64, toBlock uint64, contractWhitelist []ethCommon.Address) (filteredLogs []types.FilterParsedLog, err error) {

	if fromBlock > toBlock {
		// Just nothing
		return nil, nil
	}

	global.Logger.Debugf("Query ERC20 transfer events for blocks ( %d - %d ) ...", fromBlock, toBlock)

	var logs []logWithIA

	if len(config.Config.SendersAddress) > 0 {
		var sendersFilter []ethCommon.Hash
		for _, sender := range config.Config.SendersAddress {
			sendersFilter = append(sendersFilter, ethCommon.BytesToHash(sender.Bytes()))
		}

		sendLogs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: contractWhitelist,
			Topics: [][]ethCommon.Hash{
				{consts.ERC20_TRANSFER_TOPIC_0},
				sendersFilter,
				{},
			},
		})
		if err != nil {
			global.Logger.Errorf("Failed to query send log for blocks ( %d - %d ) with error: %v", fromBlock, toBlock, err)
		} else {
			for _, sL := range sendLogs {
				logs = append(logs, logWithIA{
					ethCommon.BytesToAddress(sL.Topics[1].Bytes()),
					sL,
				})
			}
		}
	}

	if len(config.Config.ReceiversAddress) > 0 {
		var receiversFilter []ethCommon.Hash
		for _, receiver := range config.Config.ReceiversAddress {
			receiversFilter = append(receiversFilter, ethCommon.BytesToHash(receiver.Bytes()))
		}

		receiverLogs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: contractWhitelist,
			Topics: [][]ethCommon.Hash{
				{consts.ERC20_TRANSFER_TOPIC_0},
				{},
				receiversFilter,
			},
		})
		if err != nil {
			global.Logger.Errorf("Failed to query receive log for blocks ( %d - %d ) with error: %v", fromBlock, toBlock, err)
		} else {
			for _, rL := range receiverLogs {
				logs = append(logs, logWithIA{
					ethCommon.BytesToAddress(rL.Topics[2].Bytes()),
					rL,
				})
			}
		}
	}

	blockTs := make(map[uint64]time.Time)

	// Get logs (refer to ERC20 Token Transfer Topic definitions)
	zeroHash := ethCommon.BigToHash(big.NewInt(0))
	for _, log := range logs {
		if log.L.Topics[1].Cmp(zeroHash) != 0 {
			// Not mint (mint is from null (genesis) address)
			// Set timestamp
			ts, ok := blockTs[log.L.BlockNumber]
			if !ok {
				bHeader, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(log.L.BlockNumber)))
				if err != nil {
					global.Logger.Errorf("Failed to get info for block %d with error: %v", log.L.BlockNumber, err)
					ts = time.Unix(0, 0)
				} else {
					ts = time.Unix(int64(bHeader.Time), 0)
					blockTs[log.L.BlockNumber] = ts
				}
			}
			filteredLogs = append(filteredLogs, types.FilterParsedLog{
				IAddress:    log.IA,
				BlockNumber: big.NewInt(int64(log.L.BlockNumber)),
				Sender:      ethCommon.BytesToAddress(log.L.Topics[1].Bytes()),
				Receiver:    ethCommon.BytesToAddress(log.L.Topics[2].Bytes()),
				Amount:      new(big.Int).SetBytes(log.L.Data),
				Contract:    &log.L.Address,
				TxHash:      log.L.TxHash,
				TimeStamp:   ts,
			})
		}
	}

	return filteredLogs, nil
}
