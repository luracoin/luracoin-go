package cli

import (
	"fmt"
	"github.com/maguayo/luracoin/blockchain"
	"github.com/maguayo/luracoin/wallet"
	"log"
	"os"
	"strconv"
	"time"
)

// CLI responsible for processing command line arguments
type CLI struct{}

// printUsage is similar to --help command. It prints the command line arguments
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	// REAL
	fmt.Println("  listaddresses: Lists all addresses from the wallet file")
	fmt.Println("  createwallet: Generates a new key-pair and saves it into the wallet file")

	// FAKE
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

// validateArgs function that checks if there are more than two args.
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	switch os.Args[1] {
	case "createwallet":
		createwallet()
	case "listaddresses":
		listaddresses()
	case "createGenesisBlock":
		createblockchain(os.Args[2])
	case "addBlock":
		addblock(os.Args[2])
	case "tx":
		transcation(os.Args[2])
	case "test":
		test(os.Args[2])
	case "getBlock":
		getBlock(os.Args[2])
	case "searchBlock":
		searchBlock(os.Args[2])
	default:
		cli.printUsage()
		os.Exit(1)
	}
}

func transcation(data string) {
	tx := blockchain.NewCoinbaseTX("abc", data)
	tx.PrintTransaction()

	//blockchain.NewUTXOTransaction()
}

func test(data string) {
	fmt.Println("\n===============")
	fmt.Println(time.Now().Unix())
	fmt.Println("===============\n")
	bc := blockchain.CreateGenesisBlock(data)
	fmt.Println(time.Now().Unix())
	bc.AddBlock("Hola1")
	fmt.Println(time.Now().Unix())
	bc.AddBlock("Hola2")
	fmt.Println(time.Now().Unix())
	bc.AddBlock("Hola2")
	fmt.Println(time.Now().Unix())
	bc.AddBlock("Hola3")
	fmt.Println(time.Now().Unix())
	bc.AddBlock("Hola4")
	fmt.Println(time.Now().Unix())
	bc.AddBlock("Hola5")
	fmt.Println(time.Now().Unix())
	bc.PrintBlockchain()
	fmt.Println("\n===============")
	fmt.Println(time.Now().Unix())
	fmt.Println("===============\n")
}

func createblockchain(data string) {
	blockchain.CreateGenesisBlock(data)
	fmt.Println("Done!")
}

func getBlock(hash string) {
	hash = "0000dbfe79456778b76494cbefa49534bd114c22c5fa508e31229b240182c200"
	//blockchain.SavedBlocks()
	fmt.Println("END DATABASE")
	blockchain.SearchBlookByHash(hash)
}

func searchBlock(term string) {
	a, _ := strconv.Atoi(term)
	fmt.Println(a)
	//blockchain.SearchBlook(a)
}

func addblock(data string) {
	//fmt.Println(blockchain.AddBlock(data))
	fmt.Println("Done!")
}

func createwallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}

func listaddresses() {
	wallets, err := wallet.NewWallets()
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAddresses()
	for _, address := range addresses {
		fmt.Println(address)
	}
}
