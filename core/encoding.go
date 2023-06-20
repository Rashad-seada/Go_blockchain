package core

import (
	"encoding/gob"
	"io"
)

type Encoder[T any] interface {
	Encode(T) error
}

type Decoder[T any] interface {
	Decode(T) error
}


// transaction Encoder
type GobTxEncoder struct {
	w io.Writer
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder{
	return &GobTxEncoder{
		w: w,
	}
}

func (e *GobTxEncoder) Encode(t *Transaction) error {
	return gob.NewEncoder(e.w).Encode(t)
}

// transaction Decoder
type GobTxDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder{
	return &GobTxDecoder{
		r: r,
	}
}

func (e *GobTxDecoder) Decode(t *Transaction) error {
	return gob.NewDecoder(e.r).Decode(t)
}

