// Use the main package
package main

// Create the Block data structure
type Block struct {
	Id					int64	// the id of the current block
	Timestamp			int64	// the time when the block was created
	PreviousBlockHash	[]byte	// the hash of the previous block
	MyBlockHash			[]byte 	// the hash of the current block
	AllData				[]byte	// the data or transactions (body info)
}

// Prepere the Blockchain data structure
type Blockchain struct {
	Blocks []*Block	// remember a blockchain is a series of blocks
}