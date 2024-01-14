package consts

import "github.com/ethereum/go-ethereum/crypto"

var (
	ERC20_TRANSFER_TOPIC_0 = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)
