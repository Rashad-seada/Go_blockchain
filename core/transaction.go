package core

import (
	"blockchain/types"
	"io"
)

type Transaction struct {
	data []byte
}

func (t *Transaction) Hash() types.Hash {
	
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}