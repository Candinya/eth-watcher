package types

type ContractAddressWithMeta struct {
	Address string `json:"address"`
	ContractMeta
}

type WebhookCallbackBody struct {
	ChainID     int64                    `json:"chain_id"`
	Sender      string                   `json:"sender"`
	Receiver    string                   `json:"receiver"`
	IsNative    bool                     `json:"is_native"`
	Contract    *ContractAddressWithMeta `json:"contract,omitempty"` // Only when not native
	Amount      float64                  `json:"amount"`             // Parsed amount
	Transaction string                   `json:"tx"`
}
