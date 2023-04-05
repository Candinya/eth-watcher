package config

import ethCommon "github.com/ethereum/go-ethereum/common"

type status struct {
	Receivers []ethCommon.Address
}

var Status status
