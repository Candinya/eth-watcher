package inits

import (
	"eth-watcher/config"
	"eth-watcher/global"
	"eth-watcher/jobs"
)

func Jobs() error {
	// Start watching chains
	for _, chain := range config.Config.Chain {
		err := jobs.WatchBlockChain(&chain)
		if err != nil {
			global.Logger.Errorf("Failed to start watching chain #%d with error: %v", chain.ID, err)
		}
	}

	return nil
}
