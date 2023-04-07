package jobs

import (
	"eth-watcher/config"
	"eth-watcher/global"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"testing"
)

func TestFilterERC20Transfer(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	config.Config.ReceiversHash = []ethCommon.Hash{
		ethCommon.HexToHash("0xD3E8ce4841ed658Ec8dcb99B7a74beFC377253EA"),
		ethCommon.HexToHash("0x9C8a0A9B5d5b178D73e775a2dC4D52711758C388"),
	}

	client, err := ethclient.Dial("https://rpc.sepolia.org")
	if err != nil {
		t.Fatalf("Failed to dial client with error: %v", err)
		return
	}

	// Looking for 0x9bbfcecd22e6ac1a1bbd8fa08f0f80fd12edc3ac05886da243340a26bc298f8f
	logs, err := filterERC20Transfer(client, 3229370, 3229467, nil)
	if err != nil {
		t.Fatalf("Failed to filter native transfer with error: %v", err)
		return
	}

	t.Logf("%+v\n", logs)

}
