package crypto

import (
	"blockchain/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type Singnature struct {
	r *big.Int
	s *big.Int
}

func (s *Singnature) Verify(k *ecdsa.PublicKey, data []byte) bool {
	return ecdsa.Verify(k,data,s.r,s.s)
}

type Keypair struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}

func (k *Keypair) Address() types.Address {
	h := sha256.Sum256(PublicKeyToSlice(*k.PublicKey))

	return types.NewAddressFromBytes(h[len(h)-20:])
}

func (k *Keypair) Sign(data []byte) (*Singnature,error) { 
	r , s ,err := ecdsa.Sign(rand.Reader,k.PrivateKey,data)
	if err != nil {
		return nil , err
	}
	return &Singnature{r:r,s:s} , nil
}

func GeneratePrivatekey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
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


	