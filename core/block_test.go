package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32,prevousHash types.Hash) *Block {

	k := crypto.GenerateUniqueKeypair()

	h := &Header{
		Version: 1,
		Height: height,
		PrevousHash: prevousHash,
		TimpStamp: time.Now(),
	}

	d := Data{
		Transactions: &[]Transaction{
			Transaction{
				Data: []byte("this code is written by the best coder of all times"),
			},
			Transaction{
				Data: []byte("this code is written by the best coder of all times"),
			},		
	
		},
	}

	for tx := range *d.Transactions {
		(*d.Transactions)[tx].Sign(*k)
	}

	b := NewBlock(h,d)
	b.Sign(*k)
	b.Hash(BlockHasher{})
	return b

}

func TestHashBlock(t *testing.T){
	b := randomBlock(0,types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestVerificationBlock(t *testing.T){
	keypair := crypto.GenerateUniqueKeypair()
	b := randomBlock(0,types.Hash{})

	assert.Nil(t,b.Sign(*keypair))
	assert.Nil(t,b.Verify())
}


func TestBlockSignature(t *testing.T){
	keypair := crypto.GenerateUniqueKeypair()
	b := randomBlock(0,types.Hash{})
	
	assert.Nil(t,b.Sign(*keypair))
	assert.Nil(t,b.Verify())
	assert.NotNil(t,b.Signature)

}

