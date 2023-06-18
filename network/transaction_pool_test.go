package network

import (
	"blockchain/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionPool(t *testing.T) {
	p := NewTransactionPool()
	assert.Equal(t,p.Len(),0)
}

func TestTransactionPoolAdd(t *testing.T) {
	p := NewTransactionPool()
	assert.Equal(t,p.Len(),0)

	tx1 := core.NewTransaction([]byte("my name is rashad"))
	tx2 := core.NewTransaction([]byte("my name is Zeyad"))

	assert.Nil(t,p.Add(tx1))
	assert.Nil(t,p.Add(tx2))

	assert.Equal(t,p.Len(),2)

	p.Flush()
	assert.Equal(t,p.Len(),0)
}