package network

import (
	"bytes"
	"fmt"
	"sync"
)

type LocalTransport struct {
	TransportAddress NetworkAddress
	consumeChannal chan RPC
	lock    sync.RWMutex
	peers   map[NetworkAddress]*LocalTransport
}


func NewLocalTransport(address NetworkAddress) Transport {
	return &LocalTransport{
		TransportAddress: address,
		consumeChannal: make(chan RPC,1024),
		peers: make(map[NetworkAddress]*LocalTransport),
	}
}

func (transport *LocalTransport) Consume() <-chan RPC {
	return transport.consumeChannal
}

func (localtransport *LocalTransport) Connect(transport Transport) error {
	localtransport.lock.Lock()
	defer localtransport.lock.Unlock()

	localtransport.peers[transport.Address()] = transport.(*LocalTransport)
	return nil
}

func (transport *LocalTransport) SendMessage(to NetworkAddress,payload []byte) error {
	transport.lock.RLock()
	defer transport.lock.RUnlock()

	peer , ok := transport.peers[to]
	
	if !ok {
		return fmt.Errorf("cannot send the message to this address : %s : ",to)
	}

	peer.consumeChannal <- RPC{
		from: transport.TransportAddress,
		payload: bytes.NewReader(payload),
	}

	return nil
}

func (tranport *LocalTransport) Broadcast(payload []byte) error {
	for _, peer := range tranport.peers {
		if peer.Address() != tranport.Address() {
			if err := tranport.SendMessage(peer.Address(),payload) ; err != nil{
				return err
			}
			
		}
	}

	return nil
}

func (tranport *LocalTransport) Address() NetworkAddress {
	return tranport.TransportAddress
}


