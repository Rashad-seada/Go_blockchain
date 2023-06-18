package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"bytes"
	"crypto/ecdsa"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	Data []byte
	hash types.Hash
	From *ecdsa.PublicKey
	Signature *crypto.Signature
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,

	}
}

func (t *Transaction) Bytes() []byte {
	buffer := &bytes.Buffer{}
	encoder := gob.NewEncoder(buffer)
	encoder.Encode(t.Data)
	return buffer.Bytes()
}

func (t *Transaction) Sign(k crypto.Keypair)  error {
	sig , err := k.Sign(t.Data)
	if err != nil {
		return err 
	}

	t.From = k.PublicKey
	t.Signature = sig

	return nil
}

func (t *Transaction) Verify() error {
	if t.Signature == nil {
		return fmt.Errorf("the Transaction has no signature")
	}

	if !t.Signature.Verify(t.From,t.Data) {
		return fmt.Errorf("invalid transaction signature ")
	}

	return nil
}

func (t *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if t.hash.IsZero() {
		t.hash = hasher.Hash(t)
	}

	return t.hash
}