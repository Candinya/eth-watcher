package types

type WebhookCallbackBody struct {
	ChainID     int64         `json:"chain_id"`
	Sender      string        `json:"sender"`
	Receiver    string        `json:"receiver"`
	IsNative    bool          `json:"is_native"`
	Contract    *ContractMeta `json:"contract,omitempty"` // Only when not native
	Amount      float64       `json:"amount"`             // Parsed amount
	Transaction string        `json:"tx"`
}
