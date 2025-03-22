package utils

import (
	"bytes"
	"encoding/json"
	"eth-watcher/config"
	"eth-watcher/global"
	"eth-watcher/types"
	"net/http"
	"time"
)

func WebhookCallback(chain *types.ChainConfig, ia string, sender string, receiver string, isNative bool, contractAddress *string, contractMeta *types.ContractMeta, amount float64, balance float64, tx string, ts time.Time) {
	global.Logger.Debugf("Sending response to webhooks...")

	// Prepare request body
	var body types.WebhookCallbackBody
	if isNative {
		// No contract
		body = types.WebhookCallbackBody{
			TimeStamp:   ts,
			ChainID:     chain.ID,
			IAddress:    ia,
			Sender:      sender,
			Receiver:    receiver,
			IsNative:    isNative,
			Amount:      amount,
			Balance:     balance,
			Transaction: tx,
		}
	} else {
		// With contract
		body = types.WebhookCallbackBody{
			TimeStamp: ts,
			ChainID:   chain.ID,
			IAddress:  ia,
			Sender:    sender,
			Receiver:  receiver,
			IsNative:  isNative,
			Contract: &types.ContractAddressWithMeta{
				Address:      *contractAddress,
				ContractMeta: *contractMeta,
			},
			Amount:      amount,
			Balance:     balance,
			Transaction: tx,
		}
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
			req.Header.Set("Content-Type", "application/json")
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
