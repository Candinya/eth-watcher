package config

import (
	"eth-watcher/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type config struct {
	System struct {
		Redis      string `yaml:"redis"`
		Production bool   `yaml:"production"`
	} `yaml:"system"`
	Chain    []types.ChainConfig `yaml:"chain"`
	Webhooks []string            `yaml:"webhooks"`

	ReceiversCfg  []string         `yaml:"receiver"` // Watching address config
	ReceiversHash []ethCommon.Hash `yaml:"-"`        // For ERC20 filters
}

var Config config
