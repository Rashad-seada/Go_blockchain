package core

import (
	"blockchain/types"
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeaderDecodeAndEncode(t *testing.T) {

	h := Header{
		Version: 1,
		Height: 10,
		hash: types.RandomHash(32),
		PrevousHash: types.RandomHash(32),
		TimpStamp: time.Now(),
		Nounce: 100,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t,h.EncodeBinary(buf))

	hDecoded := Header{}
	assert.Nil(t,hDecoded.DecodeBinary(buf))

	assert.Equal(t,h,hDecoded)
		
}

func TestBlockDecodeAndEncode(t *testing.T) {

	h := Header{
		Version: 1,
		Height: 10,
		hash: types.RandomHash(32),
		PrevousHash: types.RandomHash(32),
		TimpStamp: time.Now(),
		Nounce: 100,
	}

	d := Data{
		Transactions: make([]Transaction, 12),
	}

	b := Block{
		Header: h,
		Data: d,
	}

	bDecoded := Block{
		Header: h,
		Data: d,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t,b.EncodeBinary(buf))

	assert.Nil(t,bDecoded.DecodeBinary(buf))

	assert.Equal(t,b,bDecoded)
		
}

func TestBlockHash(t *testing.T){
	h := Header{
		Version: 1,
		Height: 10,
		PrevousHash: types.RandomHash(32),
		TimpStamp: time.Now(),
		Nounce: 100,
	}

	d := Data{
		Transactions: make([]Transaction, 12),
	}

	b := Block{
		Header: h,
		Data: d,
	}

	assert.NotNil(t,b.CalculateHash())

}