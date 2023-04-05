package inits

import (
	"errors"
	"eth-watcher/config"
	"fmt"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v3"
	"os"
)

func Config() error {
	// Read config file
	configFilePosition, exist := os.LookupEnv("CONFIG_FILE_PATH")
	if !exist {
		configFilePosition = "config.yml"
	}

	configFileBytes, err := os.ReadFile(configFilePosition)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFileBytes, &config.Config)
	if err != nil {
		return err
	}

	// Validate config
	for _, chain := range config.Config.Chain {
		if chain.Interval == 0 {
			// Invalid interval
			return errors.New(fmt.Sprintf("invalid block interval for blockchain #%d", chain.ID))
		}
	}

	// Update status
	for _, receiverAddress := range config.Config.Receiver {
		config.Status.Receivers = append(config.Status.Receivers, ethCommon.HexToAddress(receiverAddress))
	}

	return nil
}
