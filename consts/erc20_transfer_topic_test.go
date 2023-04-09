package consts

import (
	"bytes"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestERC20TransferTopic(t *testing.T) {
	transferEventStructure := []byte("Transfer(address,address,uint256)")
	transferEventSigHash := crypto.Keccak256Hash(transferEventStructure)

	if bytes.Equal(transferEventSigHash.Bytes(), ethCommon.HexToHash(ERC20_TRANSFER_TOPIC_0_HEX).Bytes()) {
		t.Logf("Signature matches")
	} else {
		t.Fatalf("Signature Mismatch")
	}
}
