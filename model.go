package main

import "errors"

// Node the agent hosting the blockchain networking
type Node struct{}

// Nonce the proof of work
func (n *Node) Nonce(block Block) error {
	return errors.New("Not implemented")
}

// Broadcast the block to the network
func (n *Node) Broadcast(block Block) error {
	return errors.New("Not implemented")
}

// Block the data structure representing each block in the blockchain
type Block struct {
	Hash         Hash
	PreHash      Hash
	Transactions []Transaction
	Metadata     Metadata
}

// Hash the indentifier of the block
type Hash struct {
	Nouce string
}

// Metadata the extensible data of the block
type Metadata map[string]string

// Transaction the transaction data structure
type Transaction struct{}

// Chain the chain of blocks
type Chain struct{}
