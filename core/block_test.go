package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/faridanangs/projectx/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     123456,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	hDecode.DecodeBinary(buf)
	assert.Equal(t, hDecode, h)

}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     123456,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	bDecode.DecodeBinary(buf)
	assert.Equal(t, bDecode, b)

}

func TestBlock_Hash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     123456,
		},
		Transactions: []Transaction{},
	}
	h := b.Hash()
	assert.False(t, h.IsZero())
}
