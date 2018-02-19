package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	//"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 10

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID       []byte
	Version  []byte
	Vin      []TXInput
	Vout     []TXOutput
	Locktime []byte
}

func (tx *Transaction) PrintTransaction() {
	fmt.Println("===================================================================")
	fmt.Printf("ID: %d \n", tx.ID)
	fmt.Printf("Version: %s \n", tx.Version)
	fmt.Printf("Vin: \n")
	fmt.Println("")
	for _, vin := range tx.Vin {
		fmt.Println("*******************")
		vin.PrintTXInput()
		fmt.Println("*******************")
	}
	fmt.Println("")
	fmt.Printf("Vout: \n")
	fmt.Println("")
	for _, vout := range tx.Vout {
		fmt.Println("*******************")
		vout.PrintTXOutput()
		fmt.Println("*******************")
	}
	fmt.Printf("Locktime: %d \n", tx.Locktime)
}

// IsCoinbase checks whether the transaction is coinbase. The Vin Txid has to be
// 0 and the Vin Vout has to be -1.
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

// SetID sets ID of a transaction
func (tx Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TXInput represents a transaction input There are two types of inputs.
// Coinbase / Previous transaction output. *Txid* stores the ID of such transaction
// *Vout* stores an index of an output in the transaction
// *ScriptSig* is a script which provides data to be used in an output’s ScriptPubKey
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

func (txIn *TXInput) PrintTXInput() {
	fmt.Printf("  Txid: %d \n", txIn.Txid)
	fmt.Printf("  Vout: %s \n", txIn.Vout)
	fmt.Printf("  ScriptSig: %s \n", txIn.ScriptSig)
	fmt.Println("")
}

// TXOutput represents a transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func (txOut *TXOutput) PrintTXOutput() {
	fmt.Printf("  Value: %d \n", txOut.Value)
	fmt.Printf("  ScriptPubKey: %s \n", txOut.ScriptPubKey)
	fmt.Println("")
}

// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// CanBeUnlockedWith checks if the output can be unlocked with the provided data
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, nil, []TXInput{txin}, []TXOutput{txout}, nil}
	tx.SetID()

	return &tx
}

/*
// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}
*/
