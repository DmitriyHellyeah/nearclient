package nearclient

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/nacl/sign"
)

func validatePrivateKey(key string) ([]byte, error) {
	parts := strings.Split(key, ":")
	if len(parts) == 1 {
		return base58.Decode(parts[0])
	} else if len(parts) == 2 {
		switch v := strings.ToUpper(parts[0]); v {
		case "ED25519":
			return base58.Decode(parts[1])
		default:
			return nil, fmt.Errorf("Unknown curve: %s", parts[0])
		}
	} else {
		return nil, fmt.Errorf("Invalid encoded key format, must be <curve>:<encoded key>'")
	}
}
func getKeys(key string) (publicKey *[32]byte, privateKey *[64]byte, err error) {
	validKey, err := validatePrivateKey(key)
	if err != nil {
		return nil, nil, nil
	}
	public, private, err := sign.GenerateKey(bytes.NewReader(validKey))
	return public, private, err
}
