package types

import (
	"crypto/rand"
	"encoding/hex"
)

type Hash [32]uint8

func RandomHash() Hash {
	var value [32]uint8

	token := make([]byte, 32)
	rand.Read(token)

	for i := range 32 {
		value[i] = token[i]
	}

	return Hash(value)
}

func (h *Hash) IsZero() bool {
	for i := range 32 {
		if h[i] != 0 {
			return false
		}
	}

	return true
}

func (h *Hash) ToSlice() []byte {
	b := make([]byte, 32)

	for i := range 32 {
		b[i] = h[i]
	}

	return b
}

func (h *Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}
