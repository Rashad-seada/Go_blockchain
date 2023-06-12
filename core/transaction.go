package core

import (
	"blockchain/types"
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"io"
)

type Transaction struct {
	data []byte
	hash types.Hash
}

func (t *Transaction) CalculateHash() types.Hash {
	buf := &bytes.Buffer{}
	t.EncodeBinary(buf)
	hash := sha256.Sum256(buf.Bytes())
	t.hash = hash

	return hash
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w,binary.LittleEndian, t.hash) ; err != nil {
		return nil
	}
	return binary.Write(w,binary.LittleEndian,t.data)
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r,binary.LittleEndian, &t.hash) ; err != nil {
		return nil
	}
	return binary.Read(r,binary.LittleEndian,&t.data)
}