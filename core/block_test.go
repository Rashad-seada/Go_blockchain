package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	h := &Header{
		Version: 1,
		Height: height,
		Hash: types.RandomHash(32),
		PrevousHash: types.RandomHash(32),
		TimpStamp: time.Now(),
		Nounce: 100,
	}

	d := Data{
		Transactions: make([]Transaction, 12),
	}

	return NewBlock(h,d)

}

func TestHashBlock(t *testing.T){
	b := randomBlock(0)
	fmt.Println(b.CalculateHash(BlockHasher{}))
}

func TestVerificationBlock(t *testing.T){
	keypair := crypto.GenerateUniqueKeypair()
	b := randomBlock(0)

	assert.Nil(t,b.Sign(*keypair))
	assert.Nil(t,b.Verify())
}


func TestBlockSignature(t *testing.T){
	keypair := crypto.GenerateUniqueKeypair()
	b := randomBlock(0)
	
	assert.Nil(t,b.Sign(*keypair))
	assert.Nil(t,b.Verify())
	assert.NotNil(t,b.Signature)

}

