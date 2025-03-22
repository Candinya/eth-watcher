package jobs

import (
	"eth-watcher/config"
	"eth-watcher/global"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"testing"
)

func TestFilterNativeTransferSend(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	config.Config.SendersAddress = []ethCommon.Address{
		ethCommon.HexToAddress("0xEa4d57B2dD421c5bfC893d126Ec15bc42B3d0bcD"),
	}

	client, err := ethclient.Dial("https://sepolia.drpc.org")
	if err != nil {
		t.Fatalf("Failed to dial client with error: %v", err)
		return
	}

	// Looking for 0x614fe374183bdaf259574a53e5248493a64b374b0ae13f1ee31398ca1450bc41
	logs, err := filterNativeTransfer(client, 3196892, 3196892)
	if err != nil {
		t.Fatalf("Failed to filter native transfer with error: %v", err)
		return
	}

	for _, log := range logs {
		t.Logf("%+v\n", log)
	}

}

func TestFilterNativeTransferReceive(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	config.Config.ReceiversAddress = []ethCommon.Address{
		ethCommon.HexToAddress("0xD3E8ce4841ed658Ec8dcb99B7a74beFC377253EA"),
	}

	client, err := ethclient.Dial("https://sepolia.drpc.org")
	if err != nil {
		t.Fatalf("Failed to dial client with error: %v", err)
		return
	}

	// Looking for 0x614fe374183bdaf259574a53e5248493a64b374b0ae13f1ee31398ca1450bc41
	logs, err := filterNativeTransfer(client, 3196892, 3196892)
	if err != nil {
		t.Fatalf("Failed to filter native transfer with error: %v", err)
		return
	}

	for _, log := range logs {
		t.Logf("%+v\n", log)
	}

}
