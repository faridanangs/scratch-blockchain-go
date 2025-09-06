package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) Validator {
	return &BlockValidator{bc: bc}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HashBlock(b.Height) {
		return fmt.Errorf("chain already contain block (%d) with hash (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
