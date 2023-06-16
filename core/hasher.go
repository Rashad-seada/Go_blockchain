package core

import (
	"blockchain/types"
	"crypto/sha256"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}


type TransactionHasher struct {

}

func (TransactionHasher) Hash(t *Transaction) types.Hash {
	h := sha256.Sum256(t.TransactionBytes())
	return types.Hash(h)
}


type BlockHasher struct {

}

func (BlockHasher) Hash(b *Header) types.Hash {
	h := sha256.Sum256(b.Bytes())
	return types.Hash(h)
}