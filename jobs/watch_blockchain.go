package jobs

import (
	"eth-watcher/global"
	"eth-watcher/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func WatchBlockChain(chain *types.ChainConfig) error {
	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		global.Logger.Errorf("Failed to connect to provided RPC %s for chain #%d with error: %v", chain.RPC, chain.ID, err)
		return err
	}
	global.Logger.Infof("Start watching chain #%d...", chain.ID)
	go func() {
		ticker := time.NewTicker(time.Duration(chain.Interval) * time.Second)
		for {
			select {
			case <-ticker.C:
				go routineQuery(chain, client)
			}
		}
	}()

	return nil
}
