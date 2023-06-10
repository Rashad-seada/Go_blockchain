package network

type NetworkAddress string

type RPC struct {
	from NetworkAddress
	payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetworkAddress,[]byte) error
	Address() NetworkAddress
}



