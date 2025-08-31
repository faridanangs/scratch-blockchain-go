package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic(fmt.Sprintf("given bytes with length %d should be 32", len(b)))
	}

	var value [32]uint8

	for i := range 32 {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)

	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
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

func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}
