package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
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
				Hash:      Hash{Nouce: fmt.Sprintf("%064x", 0)},
				Metadata:  Metadata{"name": "Genesis Block"},
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
	Chain   Chain
	Mempool []Transaction
}

// NewTransaction create a new transaction in current node
func (n *Node) NewTransaction(from Wallet, to Wallet, value Money) error {
	n.Mempool = append(n.Mempool, Transaction{
		From:  from,
		To:    to,
		Value: value,
	})

	return nil
}

// NewBlock create a new block in current node
func (n *Node) NewBlock(name string, metadata Metadata) Block {
	metadata["name"] = name

	return Block{
		PreHash:      n.Chain.LatestHash(),
		Transactions: n.Mempool,
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

	// simulate the proof of work, high difficulty means longer time to compute
	// when multiple nodes are competing to mine the block
	// only the node with the highest computation power will win
	hash := ""
	nonce := 0
	for !strings.HasPrefix(hash, strings.Repeat("0", n.Chain.Difficulty)) {
		nonce = nonce + 1
		hash = n.hash(value + fmt.Sprintf("%d", nonce))
	}

	block.Hash.Nouce = hash

	return block, nil
}

// hash the proof of work function
func (n *Node) hash(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Broadcast the block to the network
func (n *Node) Broadcast(block Block) error {
	return errors.New("Not implemented")
}
