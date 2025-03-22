package chain_test

import (
	"eth-watcher/chain"
	"eth-watcher/global"
	"eth-watcher/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

func TestGetERC20Balance(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	client, err := ethclient.Dial("https://sepolia.drpc.org")
	if err != nil {
		t.Fatalf("Failed to dial client with error: %v", err)
		return
	}

	balance, err := chain.GetERC20Balance(
		utils.P(ethCommon.HexToAddress("0xcb7729f2B44Ae7B86D58Bb8068f0EAD8fcF9378c")),
		utils.P(ethCommon.HexToAddress("0xD3E8ce4841ed658Ec8dcb99B7a74beFC377253EA")),
		big.NewInt(5084894),
		client,
	)

	if err != nil {
		t.Fatal(err)
	}

	if balance.Cmp(big.NewInt(0).Mul(big.NewInt(124), math.BigPow(10, 18))) != 0 {
		t.Fail()
	}

}
