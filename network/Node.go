package network

import (
	"blockchain/crypto"
	"fmt"
	"time"
)

type NodeOptions struct {
	Transports 	[]Transport
	Keypair 	crypto.Keypair
}

type Node struct {
	NodeOptions     NodeOptions
	rpcChannel  chan RPC
	quitChannel chan struct{}
}

func NewNode(options NodeOptions) *Node {
	return &Node{
		NodeOptions:     options,
		rpcChannel:  make(chan RPC),
		quitChannel: make(chan struct{}, 1),
	}
}

func (s *Node) Start() {
	s.initTransports()

	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcChannel:
			fmt.Println(string(rpc.payload))
		case <-s.quitChannel:
			break free
		case <-ticker.C:
			fmt.Println("do somthing every 5 seconds")
		}

	}
}

func (s *Node) initTransports() {
	for _, tr := range s.NodeOptions.Transports {

		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcChannel <- rpc
			}
		}(tr)

	}
}
