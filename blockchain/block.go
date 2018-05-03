package blockchain

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/fatih/color"
	"github.com/maguayo/luracoin/utils"
	"log"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp int64
	//Transactions  []*Transaction
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

func (b *Block) PrintBlock() {
	hash := hex.EncodeToString(b.Hash)
	prevHash := hex.EncodeToString(b.PrevBlockHash)
	data := string(b.Data)

	color.Yellow("\n===================================================================")
	color.Blue("-Timestamp:")
	fmt.Println(b.Timestamp)
	color.Blue("-Data:")
	fmt.Println(data)
	color.Blue("-PrevBlockHash:")
	fmt.Println(prevHash)
	color.Blue("-Hash:")
	fmt.Println(hash)
	fmt.Println(b.Hash)
	color.Blue("-Nonce:")
	fmt.Println(b.Nonce)
	color.Blue("-Height:")
	fmt.Println(b.Height)
	color.Yellow("===================================================================")
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte, height int) *Block {

	// Initialize the block with default configuration. Later we will override some
	// fields like nonce or hash.
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0, height}

	// Start working on the Proof of work, this function will return two values a
	// nonce and a hash.
	//pow := NewProofOfWork(block)
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	// Now that we have the hash and the nonce we override the default values.
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func GetBlock(hash string) *Block {
	opts := badger.DefaultOptions
	opts.Dir = "datadir/blocks"
	opts.ValueDir = "datadir/blocks"
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	block := &Block{}
	err2 := db.View(func(txn *badger.Txn) error {
		hash_bytes, err1 := hex.DecodeString(hash)
		if err1 != nil {
			log.Fatal(err1)
		}

		item, err2 := txn.Get(hash_bytes)
		if err2 != nil {
			return err2
		}

		val, err3 := item.Value()
		if err3 != nil {
			return err3
		}

		DeserializeBlock(val).PrintBlock()
		return nil
	})
	if err2 != nil {
		log.Fatal(err2)
	}
	return block
}

func SearchBlookByHeight(term int) {
	opts := badger.DefaultOptions
	opts.Dir = "datadir/blocks"
	opts.ValueDir = "datadir/blocks"
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		prefix := []byte(utils.HeigthToString(uint32(term)))
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			v, err := item.Value()
			if err != nil {
				return err
			}
			DeserializeBlock(v).PrintBlock()
		}
		return nil
	})
}

func SearchBlookByHash(hash string) {
	opts := badger.DefaultOptions
	opts.Dir = "datadir/blocks"
	opts.ValueDir = "datadir/blocks"
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	hash_bytes, err1 := hex.DecodeString(hash)
	if err1 != nil {
		log.Fatal(err1)
	}

	err2 := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		prefix := hash_bytes
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			v, err := item.Value()
			if err != nil {
				return err
			}
			DeserializeBlock(v).PrintBlock()
		}
		return nil
	})
	if err2 != nil {
		log.Fatal(err2)
	}
}

func RunProofOfWork(b *Block) int {
	//fmt.Println(b.Data)
	return 0
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Fatal(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Fatal(err)
	}

	return &block
}
