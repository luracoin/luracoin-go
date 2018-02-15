package blockchain

import (
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp int64
	//Transactions  []*Transaction
	Data          string
	PrevBlockHash string
	Hash          string
	Nonce         int
	Height        int
}

func (b *Block) PrintBlock() {
	fmt.Println("===================================================================")
	fmt.Printf("Timestamp: %d \n", b.Timestamp)
	fmt.Printf("Data: %s \n", b.Data)
	fmt.Printf("PrevBlockHash: %s \n", b.PrevBlockHash)
	fmt.Printf("Hash: %s \n", b.Hash)
	fmt.Printf("Nonce: %d \n", b.Nonce)
	fmt.Printf("Height: %d \n", b.Height)
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash string, height int) *Block {

	// Initialize the block with default configuration. Later we will override
	// some fields like nonce or hash.
	block := &Block{time.Now().Unix(), data, prevBlockHash, "", 0, height}

	// Start working on the Proof of work, this function will return two values
	// a nonce and a hash.
	//pow := NewProofOfWork(block)
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	// Now that we have the hash and the nonce we override the default values.
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func RunProofOfWork(b *Block) int {
	//fmt.Println(b.Data)
	return 0
}
