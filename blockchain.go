// Use the main package
package main

// Create the function that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	// get the index of the previous block
	previousBlockIndex := len(blockchain.Blocks) - 1
	// the previous block is needed, so let's get it
	PreviousBlock := blockchain.Blocks[previousBlockIndex]
	// create a new block containing the data and the hash of the previous block
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash, PreviousBlock.Id)
	// add that block to the chain to create a chain of blocks
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// Create the function that returns the whole blockchain and add the genesis block
// to it first
func NewBlockchain() *Blockchain {
	// create the genesis block
	genesisBlock := NewGenesisBlock()
	// the genesis block is added first to the blockchain
	blockchain := &Blockchain{[]*Block{genesisBlock}}
	// return the new blockchain
	return blockchain
}