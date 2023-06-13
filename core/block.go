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
	Transactions []Transaction
}

type Header struct {
	Version     uint32
	Height      uint32
	Hash        types.Hash
	PrevousHash types.Hash
	TimpStamp   time.Time
	Nounce      uint32
}

type Block struct {
	Header *Header
	Data   Data
	Validator *ecdsa.PublicKey
	Signature *crypto.Signature
	Hash types.Hash
}

func NewBlock(h *Header,d Data) *Block {
	return &Block{
		Header: h,
		Data: d,
	}
}

func (b *Block) Sign(keypair crypto.Keypair) error {
	sig , err := keypair.Sign(b.HeaderBytes())
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

	if !b.Signature.Verify(b.Validator,b.HeaderBytes()) {
		return fmt.Errorf("invald signature for this block")
	}

	return nil
}

func (b *Block) Decode(r io.Reader, decoder Decoder[*Block]) error {
	return decoder.Decode(r,b)
}

func (b *Block) Encode(w io.Writer, encoder Encoder[*Block]) error {
	return encoder.Encode(w,b)
}

func (b *Block) CalculateHash(hasher Hasher[*Block]) types.Hash {
	if b.Hash.IsZero() {
		b.Hash = hasher.Hash(b)
	}

	return b.Hash
}

func (b *Block) HeaderBytes() []byte {
	buffer := &bytes.Buffer{}
	encoder := gob.NewEncoder(buffer)
	encoder.Encode(b.Header)

	return buffer.Bytes()
}

