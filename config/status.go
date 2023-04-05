package config

import ethCommon "github.com/ethereum/go-ethereum/common"

type status struct {
	Receivers     []ethCommon.Address // For native filters
	ReceiversHash []ethCommon.Hash    // For ERC20 filters
}

var Status status
