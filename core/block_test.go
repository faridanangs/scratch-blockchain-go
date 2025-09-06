package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/faridanangs/projectx/crypto"
	"github.com/faridanangs/projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}

	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(h, []Transaction{tx})

}

func randomBlockWithSignature(t *testing.T, h uint32) *Block {
	privkey := crypto.GeneratePrivateKey()
	b := randomBlock(h)
	assert.Nil(t, b.Sign(privkey))

	return b
}

func TestBlock_Hash(t *testing.T) {
	b := randomBlock(9)

	fmt.Println(b.Hash(BlockHasher{}))
}

func TestBlockSign(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b := randomBlock(10)

	assert.Nil(t, b.Sign(privkey))
}

func TestBlockVerify(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b := randomBlock(10)

	assert.Nil(t, b.Sign(privkey))
	assert.Nil(t, b.Verify())

	otherPrivkey := crypto.GeneratePrivateKey()

	b.Validator = otherPrivkey.PublicKey()
	assert.NotNil(t, b.Verify())
}
