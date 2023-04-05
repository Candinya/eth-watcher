package inits

import (
	"eth-watcher/config"
	"eth-watcher/global"
	"eth-watcher/jobs"
)

func Jobs() error {
	// Start watching chains
	for index := range config.Config.Chain {
		err := jobs.WatchBlockChain(&config.Config.Chain[index])
		if err != nil {
			global.Logger.Errorf("Failed to start watching chain #%d with error: %v", config.Config.Chain[index].ID, err)
		}
	}

	return nil
}
