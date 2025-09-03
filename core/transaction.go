package core

import (
	"fmt"

	"github.com/faridanangs/projectx/crypto"
)

type Transaction struct {
	Data      []byte
	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Sign(privkey crypto.PrivateKey) error {
	sig, err := privkey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.Signature = sig
	tx.PublicKey = privkey.PublicKey()

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
