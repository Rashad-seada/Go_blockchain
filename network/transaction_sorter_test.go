package network

import (
	"blockchain/core"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionSorter(t *testing.T) {
	p := NewTransactionPool()

	txLen := 100

	for i := 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte("foo" + strconv.Itoa(i)))
		tx.SetSeen(int64(i * rand.Intn(1000)))
		assert.Nil(t,p.Add(tx))
	}

	assert.Equal(t,txLen,p.Len())

	txs := p.Transactions()

	for i := 0 ; i < len(txs) - 1 ; i++ {
		assert.True(t,txs[i].FirstSeen() < txs[1+i].FirstSeen())
	}
}