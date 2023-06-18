package network

import (
	"fmt"
	"time"
)

type NodeOptions struct {
	Transports 	[]core.Transport
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

func (n *Node) handleTransactions(tx *Transaction) error {
	if tx.Verify() {
		
	}
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
