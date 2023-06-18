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
	return types.Hash(sha256.Sum256(t.Bytes()))
}


type BlockHasher struct {

}

func (BlockHasher) Hash(b *Header) types.Hash {
	return types.Hash(sha256.Sum256(b.Bytes()))
}