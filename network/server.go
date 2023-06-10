package network

import (
	"fmt"
	"time"
)

type ServerOptions struct {
	Transports []Transport
}

type Server struct {
	options    ServerOptions
	rpcChannel chan RPC
	quitChannel chan struct{}
}

func NewServer(options ServerOptions) *Server {
	return &Server{
		options: options,
		rpcChannel: make(chan RPC),
		quitChannel: make(chan struct{},1),
	}
}


func (s *Server) Start() {
	s.initTransports()

	ticker := time.NewTicker(5*time.Second)

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

func (s *Server) initTransports() {
	for _, tr := range s.options.Transports {

		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcChannel <- rpc
			}
		}(tr)

	}
}
