package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/maguayo/luracoin/utils"
	"math"
	"math/big"
)

// 24 is an arbitrary number, our goal is to have a target that takes less than 256 bits
// in memory. And we want the difference to be significant enough, but not too big, because
// the bigger the difference the more difficult it's to find a proper hash.
// Example: 6 (zeroes) * 4(bits) = target 24
//const targetBits = 24
const targetBits = 16

var (
	maxNonce = math.MaxInt64
)

// ProofOfWork structure that holds a pointer to a block and a pointer to a target.
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork function, initializes a big.Int with the value of 1 and shift it left by
// 256 - targetBits bits. 256 is the length of a SHA-256 hash in bits, and it’s SHA-256
// hashing algorithm that we're going to use.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// prepareData in this function we just merge block fields with the target and nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run method starts looking for the proof of work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	//hex.EncodeToString(hash[:])
	return nonce, hash[:]
}
