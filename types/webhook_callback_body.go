package types

import "time"

type ContractAddressWithMeta struct {
	Address string `json:"address"`
	ContractMeta
}

type WebhookCallbackBody struct {
	TimeStamp   time.Time                `json:"ts"`
	ChainID     int64                    `json:"chain_id"`
	IAddress    string                   `json:"ia"` // Interest address
	Sender      string                   `json:"sender"`
	Receiver    string                   `json:"receiver"`
	IsNative    bool                     `json:"is_native"`
	Contract    *ContractAddressWithMeta `json:"contract,omitempty"` // Only when not native
	Amount      float64                  `json:"amount"`             // Parsed amount
	Balance     float64                  `json:"balance"`            // Parsed balance
	Transaction string                   `json:"tx"`
}
