package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/faridanangs/projectx/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Private key", privkey)

	return PrivateKey{
		key: privkey,
	}
}

func (k *PrivateKey) PublicKey() PublicKey {
	pubkey := k.key.PublicKey

	return PublicKey{
		key: &pubkey,
	}
}

func (k *PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		r: r,
		s: s,
	}, nil

}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k *PublicKey) ToSlice() []byte {
	r := elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
	fmt.Println("to Slice: ", r)
	return r
}

func (k *PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())

	fmt.Println("Address before: ", h)

	return types.GenerateAddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	s, r *big.Int
}

func (sign *Signature) Verify(data []byte, pubkey PublicKey) bool {
	return ecdsa.Verify(pubkey.key, data, sign.r, sign.s)
}
