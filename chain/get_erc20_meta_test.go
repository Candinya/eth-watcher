package chain_test

import (
	"eth-watcher/chain"
	"eth-watcher/global"
	"eth-watcher/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"testing"
)

func TestGetERC20Meta(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	client, err := ethclient.Dial("https://sepolia.drpc.org")
	if err != nil {
		t.Fatalf("Failed to dial client with error: %v", err)
		return
	}

	name, symbol, decimals, err := chain.GetERC20Meta(
		utils.P(ethCommon.HexToAddress("0xcb7729f2B44Ae7B86D58Bb8068f0EAD8fcF9378c")),
		client,
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("name: %s, symbol: %s, decimals: %d", name, symbol, decimals)

}
