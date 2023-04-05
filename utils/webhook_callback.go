package utils

import (
	"bytes"
	"encoding/json"
	"eth-watcher/config"
	"eth-watcher/global"
	"eth-watcher/types"
	"net/http"
)

func WebhookCallback(chain *types.ChainConfig, sender string, receiver string, isNative bool, contractMeta *types.ContractMeta, amount float64, tx string) {
	global.Logger.Debugf("Sending response to webhooks...")

	// Prepare request body
	body := types.WebhookCallbackBody{
		ChainID:     chain.ID,
		Sender:      sender,
		Receiver:    receiver,
		IsNative:    isNative,
		Contract:    contractMeta,
		Amount:      amount,
		Transaction: tx,
	}

	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		global.Logger.Errorf("Failed to marshal callback body %v with error: %v", body, err)
		return
	}

	for _, webhook := range config.Config.Webhooks {
		// Execute Asynchronously
		webhook := webhook
		go func() {
			// Prepare POST request with body bytes
			req, err := http.NewRequest("POST", webhook, bytes.NewReader(bodyBytes))
			if err != nil {
				global.Logger.Errorf("Failed to prepare request with error: %v", err)
				return
			}
			// Do request
			_, err = (&http.Client{}).Do(req)
			if err != nil {
				global.Logger.Errorf("Failed to do request with error: %v", err)
				return
			}
			// Ignore response
		}()

	}
}
