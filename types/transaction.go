package types

const (
	EmptyTxRoot = "11111111111111111111111111111111"
)

type TransactionStatus struct {
	Status struct {
		SuccessValue             *string `json:"SuccessValue,omitempty"`
		TransactionStatusFailure *TransactionStatusFailure `json:"Failure,omitempty"`
	} `json:"status"`
	Transaction        Transaction        `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transaction_outcome"`
	ReceiptsOutcome    []ReceiptsOutcome  `json:"receipts_outcome"`
	Receipts           []Receipt          `json:"receipts,omitempty"`
}

type TransactionStatusFailure struct {
	ActionError struct {
		Index int `json:"index"`
		Kind  struct {
			FunctionCallError struct {
				ExecutionError     string `json:"ExecutionError,omitempty"`
				MethodResolveError string `json:"MethodResolveError,omitempty"`
			} `json:"FunctionCallError"`
		} `json:"kind"`
	} `json:"ActionError"`
}

type FunctionCall struct {
	Args       string `json:"args"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
	MethodName string `json:"method_name"`
}

type Transaction struct {
	Actions    []interface{} `json:"actions"`
	Hash       string        `json:"hash"`
	Nonce      int64         `json:"nonce"`
	PublicKey  string        `json:"public_key"`
	ReceiverId string        `json:"receiver_id"`
	Signature  string        `json:"signature"`
	SignerId   string        `json:"signer_id"`
}

type TransactionOutcome struct {
	Proof     []interface{} `json:"proof"`
	BlockHash string        `json:"block_hash"`
	Id        string        `json:"id"`
	Outcome   struct {
		Logs        []interface{} `json:"logs"`
		ReceiptIds  []string      `json:"receipt_ids"`
		GasBurnt    int64         `json:"gas_burnt"`
		TokensBurnt string        `json:"tokens_burnt"`
		ExecutorId  string        `json:"executor_id"`
		Status      struct {
			SuccessReceiptId string `json:"SuccessReceiptId"`
		} `json:"status"`
	} `json:"outcome"`
}

type Receipt struct {
	PredecessorId string `json:"predecessor_id"`
	Receipt       struct {
		Action struct {
			Actions             []interface{} `json:"actions"`
			GasPrice            string        `json:"gas_price"`
			InputDataIds        []interface{} `json:"input_data_ids"`
			OutputDataReceivers []interface{} `json:"output_data_receivers"`
			SignerId            string        `json:"signer_id"`
			SignerPublicKey     string        `json:"signer_public_key"`
		} `json:"Action"`
	} `json:"receipt"`
	ReceiptId  string `json:"receipt_id"`
	ReceiverId string `json:"receiver_id"`
}

type ReceiptsOutcome struct {
	Proof []struct {
		Hash      string `json:"hash"`
		Direction string `json:"direction"`
	} `json:"proof"`
	BlockHash string `json:"block_hash"`
	Id        string `json:"id"`
	Outcome   struct {
		Logs        []interface{} `json:"logs"`
		ReceiptIds  []string      `json:"receipt_ids"`
		GasBurnt    int64         `json:"gas_burnt"`
		TokensBurnt string        `json:"tokens_burnt"`
		ExecutorId  string        `json:"executor_id"`
		Status      struct {
			SuccessValue string `json:"SuccessValue"`
		} `json:"status"`
	} `json:"outcome"`
}
