package core

import (
	"testing"

	"github.com/faridanangs/projectx/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {

	privkey := crypto.GeneratePrivateKey()
	tx := Transaction{Data: []byte("Hello world")}

	assert.Nil(t, tx.Sign(privkey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()

	tx := Transaction{Data: []byte("foo")}
	assert.Nil(t, tx.Sign(privkey))

	otherPrivkey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivkey.PublicKey()

	assert.NotNil(t, tx.Verify())

}
