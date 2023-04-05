package jobs

import (
	"context"
	"eth-watcher/consts"
	"eth-watcher/global"
	"eth-watcher/types"
	"eth-watcher/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func routineQuery(chain *types.ChainConfig, client *ethclient.Client) {
	global.Logger.Debugf("Start checking chain #%d...", chain.ID)
	currentHeight, err := client.BlockNumber(context.Background())
	if err != nil {
		global.Logger.Errorf("Failed to get current block height from block chain with error: %v", err)
		return
	}
	global.Logger.Debugf("Current block height: %d", currentHeight)

	blockHeightKey := fmt.Sprintf(consts.CACHE_KEY_BLOCK_HEIGHT, chain.ID)
	lastHeightStr, err := global.Redis.SetArgs(context.Background(), blockHeightKey, currentHeight, redis.SetArgs{
		Get: true, // Update key and return old value in one atomic operation
	}).Result()
	if err != nil {
		global.Logger.Errorf("Failed to get last block height from cache with error: %v", err)
		return
	}

	lastHeight, err := strconv.ParseUint(lastHeightStr, 10, 64)
	if err != nil {
		global.Logger.Errorf("Failed to parse last block height into uint64 with error: %v", err)
		return
	}
	global.Logger.Debugf("Last run block height: %d", lastHeight)

	// Check if is later blocks
	if lastHeight >= currentHeight {
		global.Logger.Debugf("Last height is larger, skip this run")
		return
	}

	// Filter native transfer events
	if chain.IncludeNative {
		nativeLogs, err := filterNativeTransfer(client, lastHeight, currentHeight-1)
		if err != nil {
			global.Logger.Errorf("Failed to filter native transfer logs with error: %v", err)
		}
		for _, log := range nativeLogs {
			utils.FilterCallback(chain, true, client, log)
		}
	}

	// Filter ERC20 transfer events
	erc20Logs, err := filterERC20Transfer(client, lastHeight, currentHeight-1)
	if err != nil {
		global.Logger.Errorf("Failed to filter ERC20 transfer logs with error: %v", err)
	}
	for _, log := range erc20Logs {
		utils.FilterCallback(chain, false, client, log)
	}

	global.Logger.Debugf("Routine finished for chain #%d", chain.ID)
}
