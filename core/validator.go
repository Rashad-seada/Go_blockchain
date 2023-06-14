package core

import "fmt"

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
		return fmt.Errorf("chain already contains block (%d) with hash (%s)",b.Header.Height,b.CalculateHash(BlockHasher{}))
	}

	if err := b.Verify();err != nil {
		return err
	}
	

	return nil
}
