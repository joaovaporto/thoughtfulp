// Use the main package
package main

import (
	"fmt" // the fmt package helps to print on the screen
)

func main() {
	// initialize the blockchain
	newBlockchain := NewBlockchain()

	// create 2 blocks and add 2 transactions
	newBlockchain.AddBlock("first transaction")
	newBlockchain.AddBlock("second transaction")

	// now print all the block and their contents
	for _, block := range newBlockchain.Blocks {
		fmt.Printf("Block ID: %x\n", block.Id)
		fmt.Printf("Timestamp: %x\n", block.Timestamp)
		fmt.Printf("Hash of the block: %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash)
		fmt.Printf("All the transaction: %s\n", block.AllData)
	}
}