package consts

import (
	"bytes"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestERC20TransferTopic(t *testing.T) {
	if bytes.Equal(ERC20_TRANSFER_TOPIC_0.Bytes(), ethCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef").Bytes()) {
		t.Logf("Signature matches")
	} else {
		t.Fatalf("Signature Mismatch")
	}
}
