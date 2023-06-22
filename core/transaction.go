package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	Data []byte
	From crypto.PublicKey
	Signature *crypto.Signature

	// cached version of the transaction data hash
	hash types.Hash

	// first seen is a timestamp of when this tx is first seen locally
	firstSeen int64
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

func (t *Transaction) Sign(k crypto.PrivateKey)  error {
	sig , err := k.Sign(t.Data)
	if err != nil {
		return err 
	}

	t.From = k.PublicKey()
	t.Signature = sig

	return nil
}

func (t *Transaction) Verify() error {
	if t.Signature == nil {
		return fmt.Errorf("the Transaction has no signature")
	}

	if !t.Signature.Verify(t.From ,t.Data) {
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

func (tx *Transaction) SetSeen(t int64) {
	tx.firstSeen = t
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.firstSeen
}

func (tx *Transaction) Decode( decoder Decoder[*Transaction]) error {
	gob.Register(elliptic.P256())
	return decoder.Decode(tx)
}

func (tx *Transaction) Encode(encoder Encoder[*Transaction]) error {
	gob.Register(elliptic.P256())
	return encoder.Encode(tx)
}