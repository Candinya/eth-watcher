package config

import "eth-watcher/types"

type config struct {
	System struct {
		Redis      string `yaml:"redis"`
		Production bool   `yaml:"production"`
	} `yaml:"system"`
	Receiver []string            `yaml:"receiver"` // Watching address
	Chain    []types.ChainConfig `yaml:"chain"`
	Webhooks []string            `yaml:"webhooks"`
}

var Config config
