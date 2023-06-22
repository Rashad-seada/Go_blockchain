package network

import (
	"blockchain/core"
	"blockchain/crypto"
	"time"

	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Millisecond

type NodeOptions struct {
	RPCHandler 	RPCHandler
	Transports 	[]Transport
	BlockTime 	time.Duration
	PrivateKey  *crypto.PrivateKey
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

	node :=  &Node{
		NodeOptions:    options,
		blockTime: 		options.BlockTime,
		memPool: 		NewTransactionPool(),	
		isValidator: 	options.PrivateKey != nil,
		rpcChannel: 	make(chan RPC),
		quitChannel: 	make(chan struct{}, 1),
	}

	if options.RPCHandler == nil {
		options.RPCHandler =  NewDefaultRCPHandler(node)
	}

	node.NodeOptions = options

	return node
}

func (n *Node) Start() {
	n.initTransports()
	
	ticker := time.NewTimer(n.blockTime)

	free:
		for {
			select {
				case rpc := <-n.rpcChannel:
					if err := n.NodeOptions.RPCHandler.HandleRPC(rpc) ; err != nil {
						logrus.Error(err)
					}

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

func (n *Node) ProcessTransaction(From NetworkAddress,tx *core.Transaction) error {

	hash := tx.Hash(core.TransactionHasher{})

	if n.memPool.Has(hash) {

		logrus.WithFields(logrus.Fields {
			"hash":  tx.Hash(core.TransactionHasher{}),
			},
		).Info("memory pool already has this transaction")
	
		return nil
	}

	if err := tx.Verify(); err != nil  {
		return err
	}

	tx.SetSeen(time.Now().UnixNano())
	
	logrus.WithFields(logrus.Fields {
		"hash":  tx.Hash(core.TransactionHasher{}),
		"memory pool":  n.memPool.Len(),

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
