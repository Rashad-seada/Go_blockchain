package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	Address NetworkAddress
	consumeChannal chan RPC
	lock    sync.RWMutex
	peers   map[NetworkAddress]*LocalTransport
}

func NewLocalTransport(address NetworkAddress) *LocalTransport {
	return &LocalTransport{
		Address: address,
		consumeChannal: make(chan RPC,1024),
		peers: make(map[NetworkAddress]*LocalTransport),
	}
}

func (transport *LocalTransport) consume() <-chan RPC {
	return transport.consumeChannal
}

func (localtransport *LocalTransport) connect(transport Transport) error {
	localtransport.lock.Lock()
	defer localtransport.lock.Unlock()

	localtransport.peers[transport.address()] = transport.(*LocalTransport)
	return nil
}

func (transport *LocalTransport) sendMessage(to NetworkAddress,payload []byte) error {
	transport.lock.RLock()
	defer transport.lock.RUnlock()

	peer , ok := transport.peers[to]
	if !ok {
		return fmt.Errorf("cannot send the message to this address : %s : ",to)
	}

	peer.consumeChannal <- RPC{
		from: transport.Address,
		payload: payload,
	}

	return nil
}

func (tranport *LocalTransport) address() NetworkAddress {
	return tranport.Address
}


