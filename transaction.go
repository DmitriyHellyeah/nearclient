package nearclient

import (
	"context"
	"github.com/DmitriyHellyeah/nearclient/types"
)

// SendAsyncTransaction signedTx - SignedTransaction encoded in base64
func (c *Client) SendAsyncTransaction(ctx context.Context, signedTx string) (hash string, err error) {
	err = c.RPCClient.CallWithDecode("broadcast_tx_async", []string{signedTx}, &hash)

	if err != nil {
		return
	}
	return
}

func (c *Client) SendAwaitTransaction(ctx context.Context, signedTx string) (resp types.TransactionStatus, err error) {
	err = c.RPCClient.CallWithDecode("broadcast_tx_commit", []string{signedTx}, &resp)

	if err != nil {
		return
	}
	return
}

func (c *Client) TransactionDetails(ctx context.Context, hash, sender string) (resp types.TransactionStatus, err error) {
	err = c.RPCClient.CallWithDecode("tx", []string{hash, sender}, &resp)
	if err != nil {
		return
	}

	return
}

func (c *Client) TransactionDetailsWithReceipt(ctx context.Context, hash, sender string) (resp types.TransactionStatus, err error) {
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_tx_status", []string{hash, sender}, &resp)

	if err != nil {
		return
	}

	return
}

func (c *Client) TransactionStatus(ctx context.Context, hash, sender string) (status string, txmsg string, err error) {
	response, err := c.TransactionDetails(ctx, hash, sender)
	if err != nil {
		return
	}
	status = "success"

	var Failure types.TransactionStatusFailure

	if *response.Status.TransactionStatusFailure != Failure {
		txmsg = response.Status.TransactionStatusFailure.ActionError.Kind.FunctionCallError.ExecutionError
		status = "failed"
		return
	}

	return
}

func (c *Client) IsTxFailed(ctx context.Context, hash, sender string) (res bool, txmsg string, err error) {
	status, txmsg, err := c.TransactionStatus(ctx, hash, sender)
	if err != nil {
		return false, "", err
	}

	if status == "failed" {
		return true, txmsg, err
	}

	return
}

func (c *Client) IsTxSucceeded(ctx context.Context, hash, sender string) (res bool, txmsg string, err error) {
	status, txmsg, err := c.TransactionStatus(ctx, hash, sender)
	if err != nil {
		return true, "", err
	}

	if status == "failed" {
		return false, txmsg, err
	}

	return true, "", err
}

func (c *Client) ReceiptById(ctx context.Context, receiptId string) (resp types.Receipt, err error) {
	err = c.RPCClient.CallWithDecode("EXPERIMENTAL_receipt", []string{"receipt_id", receiptId}, &resp)

	if err != nil {
		return
	}

	return
}

//func (c *Client) SendTransferTx(amount *big.Int, key, publicKey, addrFrom, addrTo string) (resp *types.TransactionStatus, err error) {
//	accessResponse, err := c.ViewAccessKey(context.Background(), addrFrom, publicKey)
//	if err != nil {
//		return nil, err
//	}
//	if permission.String != "FullAccess" {
//		return nil, fmt.Errorf("`Account %s does not have permission to send tokens using key: %s", addrFrom, string(publicKey[:]))
//	}
//	publicKeyBytes, privKeyBytes, err := getKeys(key)
//	if err != nil {
//		return nil, err
//	}
//	nonce_tx := nonce + 1
//	block_hash_dec, err := base58.Decode(block_hash)
//	if err != nil {
//		return nil, err
//	}
//	action := types.TransferAction{
//		Enum: types.TransferEnum,
//		Transfer: types.Transfer{
//			Deposit: *formatAmount(amount),
//		},
//	}
//	actions := []types.TransferAction{action}
//	block_hash_dec_fix := (*[32]byte)(block_hash_dec)
//	tx := types.TxTransfer{
//		SignerId: addrFrom,
//		PublicKey: types.PublicKey{
//			Data: *publicKeyBytes,
//		},
//		Nonce:      nonce_tx,
//		ReceiverId: addrTo,
//		Actions:    actions,
//		BlockHash:  *block_hash_dec_fix,
//	}
//	serialized_tx, err := borsh.Serialize(tx)
//	if err != nil {
//		return nil, err
//	}
//	serializedTxHash := sha256.Sum256(serialized_tx)
//	signature := sign.Sign(nil, serializedTxHash[:], privKeyBytes)
//	signature_fixed := (*[64]byte)(signature)
//	signed_tx := types.SignedTxTransfer{
//		Transaction: tx,
//		Signature: types.Signature{
//			KeyType: 0,
//			Data:    *signature_fixed,
//		},
//	}
//	data, err := borsh.Serialize(signed_tx)
//	if err != nil {
//		return nil, err
//	}
//	encoded_bs64 := base64.StdEncoding.EncodeToString(data)
//	//fmt.Println(encoded_bs64)
//	return a.SendAwaitTx(encoded_bs64)
//}
