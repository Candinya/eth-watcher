package types

import ethCommon "github.com/ethereum/go-ethereum/common"

type ChainConfig struct {
	ID            int64  `yaml:"id"`
	RPC           string `yaml:"rpc"`
	Interval      int64  `yaml:"interval"`
	IncludeNative bool   `yaml:"includeNative"`

	ContractWhitelistCfg     []string            `yaml:"contractWhitelist"` // From config
	ContractWhitelistAddress []ethCommon.Address `yaml:"-"`                 // Parsed
}
