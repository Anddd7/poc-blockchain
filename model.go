package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

const (
	GENESIS_NONCE = "000000"
)

// Node the agent hosting the blockchain networking
type Node struct {
	Chain      Chain
	LatestHash Hash
}

// NewBlock create a new block
func (n *Node) NewBlock(transactions []Transaction, metadata Metadata) Block {
	block := Block{
		PreHash:      n.LatestHash,
		Transactions: transactions,
		Metadata:     metadata,
	}
	return block
}

// Nonce the proof of work
func (n *Node) Nonce(block Block) (Block, error) {
	value, err := block.toString()
	if err != nil {
		return Block{}, err
	}
	nouce := hash(value)
	block.Hash.Nouce = nouce

	return block, nil
}

func hash(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
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

func (b *Block) toString() (string, error) {
	bytes, err := json.Marshal(b)
	return string(bytes), err
}

// Hash the indentifier of the block
type Hash struct {
	Nouce string
}

// Metadata the extensible data of the block
type Metadata map[string]string

// Transaction the transaction data structure
type Transaction struct {
	From  Wallet
	To    Wallet
	Value Money
}

// Wallet the wallet data structure
type Wallet struct {
	Address string
	Balance Money
}

// Money the money data structure
type Money struct {
	Amount int
}

// Chain the chain of blocks
type Chain struct {
	Blocks map[Hash]Block
}
