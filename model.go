package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"
)

const (
	GENESIS_NONCE = "000000"
)

// Block the data structure that represents the block in the blockchain
type Block struct {
	Hash         Hash
	PreHash      Hash
	Transactions []Transaction
	Metadata     Metadata
	Timestamp    time.Time
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

// toString capture and summarize the block data
func (b *Block) toString() (string, error) {
	bytes, err := json.Marshal(b)
	return string(bytes), err
}

// Chain the chain of blocks
type Chain struct {
	Difficulty int
	Blocks     []Block
}

// NewChain create a new chain with genesis block
func NewChain(difficulty int) Chain {
	return Chain{
		Difficulty: difficulty,
		Blocks: []Block{
			{
				Hash:     Hash{Nouce: GENESIS_NONCE},
				Metadata: Metadata{"name": "Genesis Block"},
				Timestamp: time.Now(),
			},
		},
	}
}

// AddBlock add a new block to the chain after broadcast is successful
func (c *Chain) AddBlock(block Block) {
	c.Blocks = append(c.Blocks, block)
}

// LatestHash get the latest hash of the chain
func (c *Chain) LatestHash() Hash {
	return c.Blocks[len(c.Blocks)-1].Hash
}

// Node the agent that hosts blockchain network
type Node struct {
	Chain Chain
}

// NewBlock create a new block in current node
func (n *Node) NewBlock(name string, transactions []Transaction, metadata Metadata) Block {
	metadata["name"] = name

	return Block{
		PreHash:      n.Chain.LatestHash(),
		Transactions: transactions,
		Metadata:     metadata,
		Timestamp:    time.Now(),
	}
}

// Nonce "mine" the block by finding the nonce
func (n *Node) Nonce(block Block) (Block, error) {
	value, err := block.toString()
	if err != nil {
		return Block{}, err
	}
	nouce := n.hash(value)
	block.Hash.Nouce = nouce

	return block, nil
}

// hash the proof of work function
func (n *Node) hash(value string) string {
	// use sleep to simulate the proof of work
	// high difficulty means longer time to compute
	// when multiple nodes are competing to mine the block
	// only the node with the highest computation power will win
	time.Sleep(time.Second * time.Duration(n.Chain.Difficulty))

	hasher := sha256.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Broadcast the block to the network
func (n *Node) Broadcast(block Block) error {
	return errors.New("Not implemented")
}
