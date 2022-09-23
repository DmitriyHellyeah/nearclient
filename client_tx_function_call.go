package nearclient

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/DmitriyHellyeah/nearclient/types"
	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
	"golang.org/x/crypto/nacl/sign"
)

type SignedTxFunctionCall struct {
	Transaction TxFunctionCall
	Signature   Signature
}

type TxFunctionCall struct {
	SignerId   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverId string
	BlockHash  [32]byte
	Actions    []FunctionCallAction
}

type FunctionCallAction struct {
	Enum         ActionEnum
	FunctionCall FunctionCall
}

type FunctionCall struct {
	MethodName string
	Args       []byte
	Gas        uint64
	Deposit    big.Int
}

func (a *Client) SendCallFunctionTx(ctx context.Context,methodName string, args []byte, deposit *big.Int, gas uint64, key, publicKey, addrFrom, addrTo string) (resp types.TransactionStatus, err error) {
	access_key, err := a.ViewAccessKeyFinality(ctx, addrFrom, publicKey)
	if err != nil {
		return resp, err
	}
	if access_key.Permission.String != "FullAccess" {
		return resp, fmt.Errorf("`Account %s does not have permission to send tokens using key: %s", addrFrom, string(publicKey[:]))
	}
	publicKeyBytes, privKeyBytes, err := getKeys(key)
	if err != nil {
		return resp, err
	}
	nonce_tx := uint64(access_key.Nonce + 1)
	block_hash_dec, err := base58.Decode(access_key.BlockHash)
	if err != nil {
		return resp, err
	}
	action := FunctionCallAction{
		Enum: FunctionCallEnum,
		FunctionCall:FunctionCall{
			Deposit:    *deposit,
			MethodName: methodName,
			Gas:        gas,
			Args:       args,
		},
	}
	actions := []FunctionCallAction{action}
	block_hash_dec_static := (*[32]byte)(block_hash_dec)
	tx := TxFunctionCall{
		SignerId: addrFrom,
		PublicKey: PublicKey{
			Data: *publicKeyBytes,
		},
		Nonce:      nonce_tx,
		ReceiverId: addrTo,
		Actions:    actions,
		BlockHash:  *block_hash_dec_static,
	}
	serialized_tx, err := borsh.Serialize(tx)
	if err != nil {
		return resp, err
	}
	serializedTxHash := sha256.Sum256(serialized_tx)
	signature := sign.Sign(nil, serializedTxHash[:], privKeyBytes)
	signature_fixed := (*[64]byte)(signature)
	signed_tx := SignedTxFunctionCall{
		Transaction: tx,
		Signature: Signature{
			KeyType: 0,
			Data:    *signature_fixed,
		},
	}
	data, err := borsh.Serialize(signed_tx)
	if err != nil {
		return resp, err
	}
	encoded_bs64 := base64.StdEncoding.EncodeToString(data)

    fmt.Println(encoded_bs64)
	return a.SendAwaitTransaction(ctx, encoded_bs64)
}
