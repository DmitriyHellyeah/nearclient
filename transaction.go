package nearclient

import (
	"context"
	"nearclient/types"
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

func (c *Client) TransactionDetailsWithReceipt(ctx context.Context, hash, sender string) (resp types.Receipt, err error) {
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

	if response.Status.TransactionStatusFailure != Failure {
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
