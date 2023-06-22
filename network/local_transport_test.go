package network

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T){

	traport1 := NewLocalTransport("A").(*LocalTransport)
	traport2 := NewLocalTransport("B").(*LocalTransport)

	traport1.Connect(traport2)
	traport2.Connect(traport1)


	//this is a test for local network service
	assert.Equal(t,traport1.peers[traport2.TransportAddress],traport2)
	assert.Equal(t,traport2.peers[traport1.TransportAddress],traport1)

}

func TestSendMessages(t *testing.T){

	traport1 := NewLocalTransport("A")
	traport2 := NewLocalTransport("B")

	traport1.Connect(traport2)
	traport2.Connect(traport1)


	msg := []byte("hello world")
	assert.Nil(t,traport1.SendMessage(traport2.Address(),msg))

	rpc := <- traport2.Consume()
	assert.Equal(t,rpc.payload ,bytes.NewReader(msg))
	assert.Equal(t,rpc.from, traport1.Address())
}