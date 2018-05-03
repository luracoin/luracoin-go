package blockchain

import (
	"bytes"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/maguayo/luracoin/utils"
	"log"
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

func SavedBlocks() {
	opts := badger.DefaultOptions
	opts.Dir = "datadir/blocks"
	opts.ValueDir = "datadir/blocks"

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err1 := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			v, err := item.Value()
			if err != nil {
				return err
			}
			fmt.Println("\n>>>>>>>>>>>>>>>>")
			fmt.Println(k)
			DeserializeBlock(v).PrintBlock()
			fmt.Println(">>>>>>>>>>>>>>>>")
		}
		return nil
	})
	if err1 != nil {
		log.Fatal(err1)
	}
}

func CreateGenesisBlock(data string) *Blockchain {
	newBlock := NewBlock(data, []byte{}, 0)
	bc := &Blockchain{}
	bc.blocks = append(bc.blocks, newBlock)

	fmt.Printf("///////////////\n")
	fmt.Printf("///////////////\n")
	newBlock.PrintBlock()
	fmt.Printf("///////////////\n")
	fmt.Printf("///////////////\n")

	key_db := bytes.Join(
		[][]byte{
			utils.HeigthToString(uint32(newBlock.Height)),
			newBlock.Hash,
		},
		[]byte{},
	)

	opts := badger.DefaultOptions
	opts.Dir = "datadir/blocks"
	opts.ValueDir = "datadir/blocks"

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(txn *badger.Txn) error {
		err1 := txn.Set(key_db, newBlock.Serialize())
		return err1
	})

	fmt.Println("Genesis block created")
	return bc
}
