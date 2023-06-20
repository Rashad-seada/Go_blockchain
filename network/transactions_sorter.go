package network

import (
	"blockchain/core"
	"blockchain/types"
	"sort"
)

type TransactionMapSorter struct {
	transactions []*core.Transaction
}

func NewTransactionMapSorter(transactionMap map[types.Hash]*core.Transaction) *TransactionMapSorter {
	transactions := make([]*core.Transaction, len(transactionMap))

	i := 0

	for _ , val := range transactionMap {
		transactions[i] = val
		i++
	}

	s := &TransactionMapSorter{
		transactions: transactions,
	}

	sort.Sort(s)

	return s
}

func (s TransactionMapSorter) Len() int {
	return len(s.transactions)
}

func (s TransactionMapSorter) Swap(x, y int) {
	s.transactions[x] , s.transactions[y] = s.transactions[y] , s.transactions[x]
}

func (s TransactionMapSorter) Less(x, y int) bool {
	return s.transactions[x].FirstSeen() < s.transactions[y].FirstSeen()
}