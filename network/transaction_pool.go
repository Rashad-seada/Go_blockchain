package network

import (
	"blockchain/core"
	"blockchain/types"
)



type TransactionPool struct {
	transactions map[types.Hash]*core.Transaction
}

func (tp *TransactionPool) Transactions() []*core.Transaction {
	s := NewTransactionMapSorter(tp.transactions)

	return s.transactions
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (tp *TransactionPool) Add(t *core.Transaction) error {
	hash := t.Hash(core.TransactionHasher{})
	if tp.Has(hash) {
		return nil
	}

	tp.transactions[hash] = t

	return nil
}

func (tp *TransactionPool) Has(hash types.Hash) bool {
	_ , ok := tp.transactions[hash]
	return ok
}

func (tp *TransactionPool) Len() int {
	return len(tp.transactions)
}

func (tp *TransactionPool) Flush() {
	tp.transactions = make(map[types.Hash]*core.Transaction)
}