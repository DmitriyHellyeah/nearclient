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

type SignedTxTransfer struct {
	Transaction TxTransfer
	Signature   Signature
}

type TxTransfer struct {
	SignerId   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverId string
	BlockHash  [32]byte
	Actions    []TransferAction
}

type Transfer struct {
	Deposit big.Int
}

type TransferAction struct {
	Enum     ActionEnum
	Transfer Transfer
}

func (a *Client) SendTransferTx(ctx context.Context, amount *big.Int, key, publicKey, addrFrom, addrTo string) (resp types.TransactionStatus, err error) {
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
	action := TransferAction{
		Enum: TransferEnum,
		Transfer: Transfer{
			Deposit: *amount,
		},
	}
	actions := []TransferAction{action}
	block_hash_dec_fix := (*[32]byte)(block_hash_dec)
	tx := TxTransfer{
		SignerId: addrFrom,
		PublicKey: PublicKey{
			Data: *publicKeyBytes,
		},
		Nonce:      nonce_tx,
		ReceiverId: addrTo,
		Actions:    actions,
		BlockHash:  *block_hash_dec_fix,
	}
	serialized_tx, err := borsh.Serialize(tx)
	if err != nil {
		return resp, err
	}
	serializedTxHash := sha256.Sum256(serialized_tx)
	signature := sign.Sign(nil, serializedTxHash[:], privKeyBytes)
	signature_fixed := (*[64]byte)(signature)
	signed_tx := SignedTxTransfer{
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
	return a.SendAwaitTransaction(ctx, encoded_bs64)
}
