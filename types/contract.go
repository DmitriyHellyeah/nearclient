package types

type ContractCode struct {
	CodeBase64  string `json:"code_base64"`
	Hash        string `json:"hash"`
	BlockHeight int    `json:"block_height"`
	BlockHash   string `json:"block_hash"`
}

type ContractChanges struct {
	Values []struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Proof []interface{} `json:"proof"`
	} `json:"values"`
	Proof       []interface{} `json:"proof"`
	BlockHeight int           `json:"block_height"`
	BlockHash   string        `json:"block_hash"`
}

type CallContractFunction struct {
	Result      []byte        `json:"result,omitempty"`
	Logs        []interface{} `json:"logs,omitempty"`
	Error       string        `json:"error,omitempty"`
	BlockHeight uint64        `json:"block_height,omitempty"`
	BlockHash   string        `json:"block_hash,omitempty"`
}

type ContractState struct {
	Values []struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Proof []interface{} `json:"proof"`
	} `json:"values"`
	Proof       []interface{} `json:"proof"`
	BlockHeight int           `json:"block_height"`
	BlockHash   string        `json:"block_hash"`
}

type ContractStateChanges struct {
	BlockHash string `json:"block_hash"`
	Changes   []struct {
		Cause struct {
			Type        string `json:"type"`
			ReceiptHash string `json:"receipt_hash"`
		} `json:"cause"`
		Type   string `json:"type"`
		Change struct {
			AccountId   string `json:"account_id"`
			KeyBase64   string `json:"key_base64"`
			ValueBase64 string `json:"value_base64,omitempty"`
		} `json:"change"`
	} `json:"changes"`
}
type ContractCodeChanges ContractStateChanges