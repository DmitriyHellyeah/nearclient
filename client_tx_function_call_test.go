package nearclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"
)

func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func initTestClient(t *testing.T) Client {
	client, err := NewClient(NearConfig{
		Host: "https://rpc.testnet.near.org",
	})
	if err != nil {
		t.Fatalf("Can't init test client, Error: %s", err)
	}
	return client
}

func TestClient__SendCallFunctionTx(t *testing.T) {
	type Test struct {
		isError  bool
		addrTo   string
		addrFrom string
		gas      uint64
		name     string
	}
	type Args struct {
		ReceiverId string  `json:"receiver_id"`
		Amount     string  `json:"amount"`
		Memo       *string `json:"memo"`
	}
	key := "ed25519:5XKLL4yQoBVyHCUyXrMt9898VG7My2iWomu1GC3wAW4V6eBwZGmreqpMiWfC1HiVpmAAWCe1pJ6RKNuEFgupbPjK"
	pubKey := "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X"

	amount := big.NewInt(0)
	amount.SetString("1000000000000000000000000", 10)
	amount.Mul(amount, big.NewInt(1))
	tests := []Test{
		{
			name:     "Simple data",
			addrFrom: "nexeranet.testnet",
			addrTo:   "token.arhius.testnet",
			gas:      300000000000000,
		},
	}
	args := Args{
		ReceiverId: "token.arhius.testnet",
		Amount:     amount.String(),
		Memo:       nil,
	}
	bytes, err := json.Marshal(&args)
	log.Println(args)
	if err != nil {
		t.Fatalf("JSON Marshal: %s", err.Error())
	}
	client := initTestClient(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tx, err := client.SendCallFunctionTx(context.Background(), "ft_transfer", bytes, big.NewInt(1), tt.gas, key, pubKey, tt.addrFrom, tt.addrTo)
            fmt.Println(prettyPrint(tx))
            fmt.Println("Error:",err)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if !tt.isError && tx.Transaction.Hash == "" {
				t.Fatalf("Tx didn't create, tx hash is empty")
			}
		})
	}
}
