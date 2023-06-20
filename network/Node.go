package network

import (
	"blockchain/core"
	"blockchain/crypto"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
)

var defaultBlockTime = 5 * time.Millisecond

type NodeOptions struct {
	Transports 	[]Transport
	BlockTime 	time.Duration
	Keypair 	*crypto.Keypair
}

type Node struct {
	NodeOptions     NodeOptions
	blockTime 		time.Duration
	memPool 		*TransactionPool
	isValidator 	bool
	rpcChannel  	chan RPC
	quitChannel 	chan struct{}
}

func NewNode(options NodeOptions) *Node {
	if options.BlockTime == time.Duration(0) {
		options.BlockTime = defaultBlockTime
	}

	return &Node{
		NodeOptions:    options,
		blockTime: 		options.BlockTime,
		memPool: 		NewTransactionPool(),	
		isValidator: 	options.Keypair != nil,
		rpcChannel: 	make(chan RPC),
		quitChannel: 	make(chan struct{}, 1),
	}
}

func (n *Node) Start() {
	n.initTransports()
	
	ticker := time.NewTimer(n.blockTime)

	free:
		for {
			select {
				case rpc := <-n.rpcChannel:
					fmt.Println(string(rpc.payload))

				case <-n.quitChannel:
					break free

				case <-ticker.C:
					if n.isValidator {
						n.createNewBlock()
					}
			}

		}

}

func (n *Node) createNewBlock() error {
	return nil
}

func (n *Node) handleTransactions(tx *core.Transaction) error {

	hash := tx.Hash(core.TransactionHasher{})

	if n.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields {
			"hash":  tx.Hash(core.TransactionHasher{}),
			},
		).Info("memory pool already has this transaction")
	
		return nil
	}

	if err := tx.Verify(); err == nil  {
		return err
	}
	
	logrus.WithFields(logrus.Fields {
		"hash":  tx.Hash(core.TransactionHasher{}),
		},
	).Info("adding new transaction to the memory pool")

	return n.memPool.Add(tx)
} 

func (n *Node) initTransports() {
	for _, tr := range n.NodeOptions.Transports {

		go func(tr Transport) {
			for rpc := range tr.Consume() {
				n.rpcChannel <- rpc
			}
			
		}(tr)

	}
}
