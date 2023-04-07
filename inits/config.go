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

	// Set receivers
	for _, receiverAddressHex := range config.Config.ReceiversCfg {
		config.Config.ReceiversHash = append(config.Config.ReceiversHash, ethCommon.HexToHash(receiverAddressHex))
	}

	// Set chain whitelists
	for index := range config.Config.Chain {
		for _, contractAddressHex := range config.Config.Chain[index].ContractWhitelistCfg {
			config.Config.Chain[index].ContractWhitelistAddress = append(config.Config.Chain[index].ContractWhitelistAddress, ethCommon.HexToAddress(contractAddressHex))
		}
	}

	return nil
}
