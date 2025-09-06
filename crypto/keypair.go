package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

type PrivateKey struct {
	key ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{key: *privkey}
}

func (priv PrivateKey) PublicKey() PublicKey {
	pubKey := priv.key.PublicKey

	return PublicKey{key: pubKey}
}

func (priv PrivateKey) Sign(data []byte) (*Signature, error) {
	// priv.key.Sign(rand.Reader, data, )

	return nil, nil
}

type PublicKey struct {
	key ecdsa.PublicKey
}

type Signature struct {
	s, r *big.Int
}
