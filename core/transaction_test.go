package core

import (
	"blockchain/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {

	Keypair := crypto.GenerateUniqueKeypair()
	tx1 := Transaction{
		Data: []byte("this code is written by the best coder of all times"),
	}

	assert.Nil(t,tx1.Sign(*Keypair))
	assert.NotNil(t,tx1.Signature)
}

func TestTransactionSignature(t *testing.T) {
	Keypair := crypto.GenerateUniqueKeypair()
	tx1 := Transaction{
		Data: []byte("this code is written by the best coder of all times"),
	}

	assert.Nil(t,tx1.Sign(*Keypair))
	assert.NotNil(t,tx1.Signature)
}

func TestTransactionVerfication(t *testing.T) {
	
	Keypair1 := crypto.GenerateUniqueKeypair()
	Keypair2 := crypto.GenerateUniqueKeypair()

	tx1 := Transaction{
		Data: []byte("this code is written by the best coder of all times"),
	}

	assert.Nil(t,tx1.Sign(*Keypair1))
	assert.Nil(t,tx1.Verify())

	assert.NotNil(t,tx1.Signature.Verify(Keypair2.PublicKey,tx1.Data))

	tx1.Data = []byte("rashad")
	assert.NotNil(t,tx1.Verify())

}
	


