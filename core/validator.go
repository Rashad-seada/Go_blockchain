package core

import (
	"fmt"
)

type Validator[T any] interface {
	ValidateBlock(*T) error
}

type BlockValidator struct {
	chain *Blockchain
}

func NewBlockValidator(chain *Blockchain) Validator[Block] {
	return &BlockValidator{
		chain: chain,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {

	if v.chain.HasBlock(b.Header.Height) {
		return fmt.Errorf("chain already contains block (%d) with hash (%s)", b.Header.Height, b.Hash(BlockHasher{}))
	}

	if b.Header.Height != v.chain.Height()+1 {
		return fmt.Errorf("Block you tried adding have invalid height")
	}

	prevousHeader, err := v.chain.GetHeader(v.chain.Height())
	if err != nil {
		return err
	}

	prevousBlockHash := BlockHasher{}.Hash(prevousHeader)
	if prevousBlockHash != b.Header.PrevousHash {
		return fmt.Errorf("the calculated hash of the prevous block is (%s) is not equal to prevous hash property (%d)", prevousBlockHash, prevousHeader.hash)
	}


	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
