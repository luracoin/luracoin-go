package blockchain

import (
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Height+1)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) GetHeight() int {
	latestBlock := bc.blocks[len(bc.blocks)]
	return latestBlock.Height
}

func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.blocks {
		block.PrintBlock()
	}
}

func CreateGenesisBlock(data string) *Blockchain {
	newBlock := NewBlock(data, "0", 0)
	bc := &Blockchain{}
	bc.blocks = append(bc.blocks, newBlock)
	fmt.Println("Genesis block created")
	return bc
}
