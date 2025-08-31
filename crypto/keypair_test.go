package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignVeirfySuccess(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubkey := privkey.PublicKey()
	addr := pubkey.Address()
	fmt.Println("Address test", addr)

	msg := []byte("Hello")
	sign, err := privkey.Sign(msg)
	assert.Nil(t, err)

	fmt.Println("Signature", sign)

	assert.True(t, sign.Verify(msg, pubkey))
}

func TestSignVeirfyFailed(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubkey := privkey.PublicKey()

	msg := []byte("Hello")
	sign, err := privkey.Sign(msg)
	assert.Nil(t, err)
	assert.False(t, sign.Verify([]byte("Hello World"), pubkey))

	privkey2 := GeneratePrivateKey()
	pubkey2 := privkey2.PublicKey()

	assert.False(t, sign.Verify(msg, pubkey2))

}
