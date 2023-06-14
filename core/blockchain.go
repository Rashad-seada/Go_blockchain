package core 

type Blockchain struct {
	Storage Storage
	Headers []*Header
	Validator Validator[Block]
}

func NewBlockchain(genesisBlock *Block) (*Blockchain,error) {
	bc := &Blockchain{
		Headers: []*Header{},
		Storage: NewMemoryStorage(),
	}
	bc.Validator = NewBlockValidator(bc)

	bc.addBlockWithoutValidation(genesisBlock)
	return bc , nil
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.Validator.ValidateBlock(b) ; err != nil {
		return err
	}
	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.Headers) - 1) 
}

func (bc *Blockchain) SetValidator(v Validator[Block])  {
	bc.Validator = v
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) createGenesisBlock(){
	
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.Headers = append(bc.Headers,b.Header )
	return bc.Storage.Put(b)
}