package types

type ChainConfig struct {
	ID            int64  `yaml:"id"`
	RPC           string `yaml:"rpc"`
	Interval      int64  `yaml:"interval"`
	IncludeNative bool   `yaml:"includeNative"`
}
