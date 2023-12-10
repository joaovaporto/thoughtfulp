// Use the main package
package main

import (
	"bytes"			// need to convert data into byte in order to be sent on the network
	"crypto/sha256"	// crypto libray to hash the data
	"strconv"		// for conversion
	"time"			// the time for out timestamp
)

// Now let's create a function for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	// get the id and convert it into a unique series of digits
	blockId := []byte(strconv.FormatInt(block.Id, 10))
	// get the time and convert it into a unique series of digits
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	// concatenate all the block data
	headers := bytes.Join([][]byte{blockId, timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	// hash the whole thing
	hash := sha256.Sum256(headers)
	// now set the hash of the block
	block.MyBlockHash = hash[:]
}

// Create a function for new block generation and return that block
func NewBlock(data string, prevBlockHash []byte, previousBlockId int64) *Block {
	// calculate the new block id
	blockId := previousBlockId + 1
	// the block is received
	block := &Block{blockId, time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	// the block is hashed
	block.SetHash()
	// the block is returned with all the information in it
	return block
}

// Let's now create the genesis block function that will return the first block.
// The genesis block is the first ever mined block, i.e. the first block on the chain
func NewGenesisBlock() *Block {
	// set the genesis block id
	genesisBlockId := int64(0)
	// the genesis block is made with some data in it
	return NewBlock("Genesis Block", []byte{}, genesisBlockId)
}