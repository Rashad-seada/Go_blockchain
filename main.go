package main

import "blockchain/network"

func main() {
	localTransport := network.NewLocalTransport("LOCAL")
	remoteTransport := network.NewLocalTransport("REMOTE")


	localTransport.Connect(remoteTransport)
	remoteTransport.Connect(localTransport)

	msg := []byte("hello local data")

	go func ()  {
		remoteTransport.SendMessage(localTransport.Address(),msg)
	}()

	options := network.ServerOptions{
		Transports: []network.Transport{localTransport},
	}

	server := network.NewServer(options)
	server.Start()

}


