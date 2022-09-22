package nearclient

import (
	"time"

	"github.com/near/borsh-go"
)

type ActionEnum borsh.Enum

const (
	CreateAccountEnum ActionEnum = iota
	DeployContractEnum
	FunctionCallEnum
	TransferEnum
	StakeEnum
	AddKeyEnum
	DeleteKeyEnum
	DeleteAccountEnum
)

type Signature struct {
	KeyType uint8
	Data    [64]byte
}

type PublicKey struct {
	KeyType uint8
	Data    [32]byte
}

type NearConfig struct {
	Host           string
	WaitingTimeout time.Duration
}
