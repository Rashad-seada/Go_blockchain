package main

import (
	"blockchain/network"
	"time"
)

func main() {
	localTransport := network.NewLocalTransport("LOCAL")
	remoteTransport := network.NewLocalTransport("REMOTE")

	localTransport.Connect(remoteTransport)
	remoteTransport.Connect(localTransport)

	msg := []byte("hello local data")

	go func() {
		for {
			time.Sleep(1 * time.Second)
			remoteTransport.SendMessage(localTransport.Address(), msg)	
		}
	}()

	options := network.NodeOptions{
		Transports: []network.Transport{localTransport},
	}

	Node := network.NewNode(options)
	Node.Start()

}
