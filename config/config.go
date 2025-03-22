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

	SendersCfg     []string            `yaml:"sender"`
	SendersAddress []ethCommon.Address `yaml:"-"`

	ReceiversCfg     []string            `yaml:"receiver"` // Watching address config
	ReceiversAddress []ethCommon.Address `yaml:"-"`        // For ERC20 filters
}

var Config config
