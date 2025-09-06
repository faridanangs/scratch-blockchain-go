package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlock := 1000

	for i := range lenBlock {
		err := bc.AddBlock(randomBlockWithSignature(t, uint32(i+1)))
		assert.Nil(t, err)
	}

	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock+1)
	assert.Nil(t, bc.AddBlock(randomBlockWithSignature(t, 1002)))
}

func TestHashBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HashBlock(uint32(0)))
	assert.Equal(t, bc.Height(), uint32(0))
}
