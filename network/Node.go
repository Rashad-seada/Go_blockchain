package network

import (
	"blockchain/core"
	"blockchain/crypto"
	"time"

	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Millisecond

type NodeOptions struct {
	RPCDecodeFunc 	RPCDecodeFunc
	RPCProcessor 	RPCProcessor
	Transports 		[]Transport
	BlockTime 		time.Duration
	PrivateKey  	*crypto.PrivateKey
}

type Node struct {
	NodeOptions     NodeOptions
	memPool 		*TransactionPool
	isValidator 	bool
	rpcChannel  	chan RPC
	quitChannel 	chan struct{}
}

func NewNode(options NodeOptions) *Node {


	if options.BlockTime == time.Duration(0) {
		options.BlockTime = defaultBlockTime
	}

	if options.RPCDecodeFunc == nil {
		options.RPCDecodeFunc =  DefaultRPCDecodeFunc
	}

	node :=  &Node{
		NodeOptions:    options,
		memPool: 		NewTransactionPool(),	
		isValidator: 	options.PrivateKey != nil,
		rpcChannel: 	make(chan RPC),
		quitChannel: 	make(chan struct{}, 1),
	}

	if options.RPCProcessor == nil {
		node.NodeOptions.RPCProcessor = node
	}
	
	if node.isValidator {
		go node.ValidatorLoop()
	}
	
	return node
}

func (n *Node) Start() {
	n.initTransports()
	
	free:
	
		for {

			select {
				case rpc := <-n.rpcChannel:
					msg , err := n.NodeOptions.RPCDecodeFunc(rpc)
					if err != nil {
						logrus.Error(err)
					}

					if err := n.NodeOptions.RPCProcessor.ProcessMessaage(msg) ; err != nil {
						logrus.Error(err)
					}

				case <-n.quitChannel:
					break free
			}

		}

}

func (n *Node) createNewBlock() error {
	return nil
}

func (n *Node) ValidatorLoop() {
	ticker := time.NewTicker(n.NodeOptions.BlockTime)

	for {
		<- ticker.C
		n.createNewBlock()
	}
}

func (n *Node) ProcessMessaage(msg *DecodedMessage) error {
	switch t := msg.Data.(type) {
	case *core.Transaction :
		return n.processTransaction(t)
	}

	return nil
}

func (n *Node) processTransaction(tx *core.Transaction) error {

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
