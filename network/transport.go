package network

type NetworkAddress string

type RPC struct {
	from NetworkAddress
	payload []byte
}

type Transport interface {
	consume() <-chan RPC
	connect(Transport) error
	sendMessage(NetworkAddress,[]byte) error
	address() NetworkAddress
}



