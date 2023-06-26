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

type DecodedMessage struct {
	From NetworkAddress
	Data any
}

type RPCDecodeFunc func (RPC) (*DecodedMessage,error)

func DefaultRPCDecodeFunc(rpc RPC) (*DecodedMessage,error) {
	msg := Message{}
	if err := gob.NewDecoder(rpc.payload).Decode(&msg) ; err != nil {
		return nil , fmt.Errorf("failed to decode message from this node %s : %s",rpc.from,err)
	}

	switch msg.Header {

	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))) ; err != nil {
			return nil , err
		}
		return &DecodedMessage{
			From: rpc.from,
			Data: tx,
		} , nil


	default: 
		return nil , fmt.Errorf("invalid message header %x",msg.Header)
	}
}


type RPCProcessor interface {
	ProcessMessaage(*DecodedMessage) error
}