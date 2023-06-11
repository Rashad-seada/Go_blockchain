package core

import (
	"blockchain/types"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"
)

type Data struct {
	Transactions []Transaction
}

func (d *Data) EncodeBinary(w io.Writer) error {
	for _ , tx := range d.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (d *Data) DecodeBinary(r io.Reader) error {
	for _ , tx := range d.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil
}

type Header struct {
	Version     uint32
	Height      uint32
	hash        types.Hash
	PrevousHash types.Hash
	TimpStamp   uint64
	Nounce      uint32
}

func (h *Header) Hash() types.Hash {
	buf := &bytes.Buffer{}
	h.EncodeBinary(buf)

	hash := types.Hash(sha256.Sum256(buf.Bytes()))
	return hash
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w,binary.LittleEndian,h.Version) ; err != nil {
		return err
	}
	if err := binary.Write(w,binary.LittleEndian,h.TimpStamp) ; err != nil {
		return err
	}
	if err := binary.Write(w,binary.LittleEndian,h.PrevousHash) ; err != nil {
		return err
	}
	if err := binary.Write(w,binary.LittleEndian,h.Hash) ; err != nil {
		return err
	}
	if err := binary.Write(w,binary.LittleEndian,h.Height) ; err != nil {
		return err
	}
	return binary.Write(w,binary.LittleEndian,h.Nounce) 

}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r,binary.LittleEndian,&h.Version) ; err != nil {
		return err
	}
	if err := binary.Read(r,binary.LittleEndian,&h.TimpStamp) ; err != nil {
		return err
	}
	if err := binary.Read(r,binary.LittleEndian,&h.PrevousHash) ; err != nil {
		return err
	}
	if err := binary.Read(r,binary.LittleEndian,&h.hash) ; err != nil {
		return err
	}
	if err := binary.Read(r,binary.LittleEndian,&h.Height) ; err != nil {
		return err
	}
	return binary.Read(r,binary.LittleEndian,&h.Nounce) 

}

type Block struct {
	Header Header
	Data   Data
}

func (b *Block) CalculateHash() types.Hash{
		buf := &bytes.Buffer{}
		b.Header.EncodeBinary(buf)
		b.Data.EncodeBinary(buf)
		hash := sha256.Sum256(buf.Bytes())
		return types.Hash(hash)
}

func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}

	if err := b.Data.EncodeBinary(w); err != nil {
		return err
	}
 
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	if err := b.Data.DecodeBinary(r); err != nil {
		return err
	}
 
	return nil
}	