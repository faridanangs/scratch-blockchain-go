package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubkey := privkey.PublicKey()
	address := pubkey.Address()

	msg := []byte("Hello cui")
	sign, err := privkey.Sign(msg)
	assert.Nil(t, err)

	b := sign.Verify(pubkey.key, msg, sign.r, sign.s)
	assert.True(t, b)

	fmt.Println(address)

}
