package core

import (
	"blockchain/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	bc , err := NewBlockchain(randomBlock(1))
	assert.Nil(t,err)
	assert.NotNil(t,bc)
	assert.Equal(t,bc.Height(),uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc , err := NewBlockchain(randomBlock(1))
	assert.Nil(t,err)
	assert.True(t,bc.HasBlock(0))

}
 
func TestAddBlock(t *testing.T) {
	bc , err := NewBlockchain(randomBlock(0))
	assert.Nil(t,err)

	bcLen := 1000

	for i := 0 ; i < bcLen ;i++ {
		b := randomBlock(uint32(i + 1))

		b.Sign(*crypto.GenerateUniqueKeypair())

		assert.Nil(t,bc.AddBlock(b))
	}


	
	assert.Equal(t,bc.Height(), uint32(bcLen))
	assert.Equal(t,len(bc.Headers), (bcLen + 1))

	assert.NotNil(t,bc.AddBlock(randomBlock(100)))

}