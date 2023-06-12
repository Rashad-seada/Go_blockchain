package core

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionEncodingAndDecoding(t *testing.T) {

	tx1 := Transaction{
		data: make([]byte, 12),
	}
	tx1.CalculateHash()

	buf := &bytes.Buffer{}

	tx2 := Transaction{}

	assert.Nil(t,tx1.EncodeBinary(buf))

	assert.Nil(t,tx2.DecodeBinary(buf))

	assert.Equal(t,tx1,tx2.DecodeBinary(buf))

}