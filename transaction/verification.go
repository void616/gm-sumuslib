package transaction

import (
	"fmt"

	sumuslib "github.com/void616/gm-sumuslib"
	"github.com/void616/gm-sumuslib/signer"
	"golang.org/x/crypto/sha3"
)

// Verify transaction payload
func Verify(from sumuslib.PublicKey, payload []byte, signature sumuslib.Signature) error {
	if len(payload) == 0 {
		return fmt.Errorf("invalid payload")
	}

	// make payload digest
	hasher := sha3.New256()
	_, err := hasher.Write(payload)
	if err != nil {
		return err
	}
	digest := hasher.Sum(nil)

	// verify
	return signer.Verify(from, digest, signature)
}
