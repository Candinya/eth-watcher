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
)

var (
	erc20TransferTopic0Hash ethCommon.Hash
)

func init() {
	erc20TransferTopic0Hash = ethCommon.HexToHash(consts.ERC20_TRANSFER_TOPIC_0_HEX)
}

func filterERC20Transfer(client *ethclient.Client, fromBlock uint64, toBlock uint64) (filteredLogs []types.FilterParsedLog, err error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Topics: [][]ethCommon.Hash{
			{erc20TransferTopic0Hash},
			{},                          // From any
			config.Status.ReceiversHash, // To receivers
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		global.Logger.Errorf("Failed to filter log for block ( %d - %d ) with error: %v", fromBlock, toBlock, err)
		return nil, err
	}

	// Get logs (refer to ERC20 Token Transfer Topic definitions)
	for _, log := range logs {
		if big.NewInt(0).Cmp(log.Topics[1].Big()) != 0 {
			// Not mint (mint is from null address)
			filteredLogs = append(filteredLogs, types.FilterParsedLog{
				Sender:   ethCommon.BytesToAddress(log.Topics[1].Bytes()),
				Receiver: ethCommon.BytesToAddress(log.Topics[2].Bytes()),
				Amount:   new(big.Int).SetBytes(log.Data),
				Contract: &log.Address,
				TxHash:   log.TxHash,
			})
		}
	}

	return filteredLogs, nil
}
