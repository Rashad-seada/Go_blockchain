package network

import (
	"blockchain/core"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MessageTypeBlock MessageType = 0x2
)

type RPC struct {
	from   	NetworkAddress
	payload io.Reader
}

type Message struct {
	Header MessageType
	Data []byte
}

func NewMessage(t MessageType,data []byte) *Message {
	return &Message{
		Header: t,
		Data: data,
	}
}

func (m *Message) Bytes() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(m)
	return buf.Bytes()
}
 
type RPCHandler interface {
	HandleRPC(rpc RPC) error
}

type DefaultRCPHandler struct {
	p RPCProcessor
}

func NewDefaultRCPHandler(p RPCProcessor) *DefaultRCPHandler{
	return &DefaultRCPHandler{
		p:p,
	}
}

func (h *DefaultRCPHandler) HandleRPC(rpc RPC) error {

	msg := Message{}
	if err := gob.NewDecoder(rpc.payload).Decode(&msg) ; err != nil {
		return fmt.Errorf("Failed to decode message from this node %s : %s",rpc.from,err)
	}

	switch msg.Header {

	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))) ; err != nil {
			return err
		}
		return h.p.ProcessTransaction(rpc.from,tx)


	default: 
		return fmt.Errorf("invalid message header %x",msg.Header)

		
	}

}


type RPCProcessor interface {
	ProcessTransaction(NetworkAddress, *core.Transaction) error
}