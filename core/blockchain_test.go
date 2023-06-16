package core

import (
	"blockchain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(1, types.Hash{}))
	assert.Nil(t, err)
	assert.NotNil(t, bc)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(1, types.Hash{}))
	assert.Nil(t, err)
	assert.True(t, bc.HasBlock(0))

}

func TestGetHeader(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)

	len := 5

	for i := 0; i < len; i++ {
		prevousHeader, err := bc.GetHeader(bc.Height())
		assert.Nil(t, err)

		b := randomBlock(uint32(i+1), prevousHeader.hash)

		assert.Nil(t, bc.AddBlock(b))

		header, err := bc.GetHeader(uint32(i + 1))

		assert.Nil(t, err)
		assert.Equal(t, header, b.Header)
	}
}

func TestAddBlock(t *testing.T) {
	b1 := randomBlock(0, types.RandomHash(32))
	bc, err := NewBlockchain(b1)

	assert.Nil(t, err)
	b2 := randomBlock(1, b1.Hash(BlockHasher{}))

	assert.Nil(t, bc.AddBlock(b2))

}

