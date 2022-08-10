package types

type Account struct {
	Amount        string `json:"amount"`
	Locked        string `json:"locked"`
	CodeHash      string `json:"code_hash"`
	StorageUsage  int    `json:"storage_usage"`
	StoragePaidAt int    `json:"storage_paid_at"`
	BlockHeight   int    `json:"block_height"`
	BlockHash     string `json:"block_hash"`
}

type AccountChanges struct {
	BlockHash string        `json:"block_hash"`
	Changes   []struct {
		Cause struct {
			Type        string `json:"type"`
			TxHash      string `json:"tx_hash,omitempty"`
			ReceiptHash string `json:"receipt_hash,omitempty"`
		} `json:"cause"`
		Type   string `json:"type"`
		Change struct {
			AccountId     string `json:"account_id"`
			Amount        string `json:"amount"`
			Locked        string `json:"locked"`
			CodeHash      string `json:"code_hash"`
			StorageUsage  int    `json:"storage_usage"`
			StoragePaidAt int    `json:"storage_paid_at"`
		} `json:"change"`
	} `json:"changes"`
}
