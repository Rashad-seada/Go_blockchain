package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"bytes"
	"crypto/ecdsa"
	"encoding/gob"
	"fmt"
	"io"
	"time"
)

type Data struct {
	Transactions *[]Transaction
}

type Header struct {
	Version     uint32
	Height      uint32
	hash        types.Hash
	PrevousHash types.Hash
	TimpStamp   time.Time
}

func (h *Header) Bytes() []byte {
	buffer := &bytes.Buffer{}
	encoder := 	gob.NewEncoder(buffer)
	encoder.Encode(h.Version)
	encoder.Encode(h.Height)
	encoder.Encode(h.PrevousHash)
	encoder.Encode(h.TimpStamp)
	return buffer.Bytes()
}

type Block struct {
	Header        *Header
	Data          Data
	Validator     *ecdsa.PublicKey
	Signature     *crypto.Signature

	//chached version of the header
	DataHash          types.Hash
}

func NewBlock(h *Header, d Data) *Block {
	return &Block{
		Header: h,
		Data:   d,
	}
}

func (b *Block) AddTransaction(transaction *Transaction) {
	*b.Data.Transactions = append(*b.Data.Transactions,*transaction)
}

func (b *Block) Sign(keypair crypto.Keypair) error {
	sig, err := keypair.Sign(b.Header.Bytes())
	if err != nil {
		return err
	}

	b.Signature = sig
	b.Validator = keypair.PublicKey
	return nil
}

func (b *Block) Verify() error {

	if b.Signature == nil {
		return fmt.Errorf("there is no signature for this block")
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("invald signature for this block")
	}

	for _, tx := range *b.Data.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) Decode(r io.Reader, decoder Decoder[*Block]) error {
	return decoder.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, encoder Encoder[*Block]) error {
	return encoder.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.DataHash.IsZero() {
		hash := hasher.Hash(b.Header)
		b.DataHash = hash
		b.Header.hash = hash

	}

	return b.DataHash
}
