package main

import (
	"blockchain/core"
	"blockchain/crypto"
	"blockchain/network"
	"github.com/sirupsen/logrus"
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	localTransport := network.NewLocalTransport("LOCAL")
	remoteTransport := network.NewLocalTransport("REMOTE")

	localTransport.Connect(remoteTransport)
	remoteTransport.Connect(localTransport)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			if err := sendTransaction(localTransport,remoteTransport.Address()); err != nil {
				logrus.Error(err)
			}
		}
	}()

	options := network.NodeOptions{
		Transports: []network.Transport{remoteTransport},
	}

	Node := network.NewNode(options)
	Node.Start()

}

func sendTransaction(transport network.Transport,to network.NetworkAddress) error {
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000)),10))
	
	tr := core.NewTransaction(data)
	
	tr.Sign(crypto.GeneratePrivateKey())

	tr.Hash(core.TransactionHasher{})

	buffer := &bytes.Buffer{}
	if err := tr.Encode(core.NewGobTxEncoder(buffer)) ; err != nil{
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx,buffer.Bytes())

	return transport.SendMessage(to,msg.Bytes())
}
