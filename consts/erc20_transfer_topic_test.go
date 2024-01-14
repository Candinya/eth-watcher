package consts

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestERC20TransferTopic(t *testing.T) {
	if ERC20_TRANSFER_TOPIC_0.Cmp(ethCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")) == 0 {
		t.Logf("Signature matches")
	} else {
		t.Fatalf("Signature Mismatch")
	}
}
