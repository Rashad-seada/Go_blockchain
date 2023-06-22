package network

type NetworkAddress string

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetworkAddress,[]byte) error
	Address() NetworkAddress
}



