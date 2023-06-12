package crypto

import (
	"blockchain/types"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/ecdsa"

)

type Keypair struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}


func GeneratePrivatekey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func GeneratePublickey(k ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &k.PublicKey
}

func GenerateUniqueKeypair() *Keypair {
	privateKey := GeneratePrivatekey()
	publicKey := GeneratePublickey(*privateKey)
	keypair := &Keypair{
		PrivateKey: privateKey,
		PublicKey: publicKey,
	}
	return keypair
}

func PublicKeyToSlice(k ecdsa.PublicKey) []byte {
	return elliptic.MarshalCompressed(k,k.X,k.Y)
}

func (k *Keypair) Address() types.Address {
	h := sha256.Sum256(PublicKeyToSlice(*k.PublicKey))

	return types.NewAddressFromBytes(h[len(h)-20:])
}
	