package core

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	Storage   	Storage
	lock 		sync.RWMutex
	Headers   	[]*Header
	Validator 	Validator[Block]
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		Headers: []*Header{},
		Storage: NewMemoryStorage(),
	}
	bc.Validator = NewBlockValidator(bc)

	bc.addBlockWithoutValidation(genesis)
	return bc, nil
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.Validator.ValidateBlock(b); err != nil {
		return err
	}
	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) Height() uint32 {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return uint32(len(bc.Headers) - 1)
}

func (bc *Blockchain) SetValidator(v Validator[Block]) {
	bc.Validator = v
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height does not exists")
	}
	bc.lock.Lock()
	defer bc.lock.Unlock()
	
	return bc.Headers[height], nil
}

// func (bc *Blockchain) createGenesisBlock() *Block {
// 	key := crypto.GenerateUniqueKeypair()
// 	b := &Block{
// 		Header: &Header{
// 			Version: 1,
// 			Height: 0,
// 			TimpStamp: time.Now(),
// 			PrevousHash: types.Hash{},
// 		},
// 		Data: Data{},
// 		Signature: nil,
// 	}
// 	b.Sign(*key)
// 	b.Validator = key.PublicKey
// 	b.CalculateHash(BlockHasher{})

// 	return b
// }

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.lock.Lock() 
	bc.Headers = append(bc.Headers, b.Header)
	bc.lock.Unlock()

	logrus.WithFields(
		logrus.Fields{
			"height": b.Header.hash,
			"hash":   b.Hash(BlockHasher{}),
		},
	).Info("")
	return bc.Storage.Put(b)
}
