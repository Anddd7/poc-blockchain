package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Node the agent that hosts blockchain network
type Node struct {
	Chain   Chain
	Mempool []Transaction
}

// SubmitTransaction create a new transaction in current node
func (n *Node) SubmitTransaction(from Wallet, to Wallet, value Money) error {
	n.Mempool = append(n.Mempool, Transaction{
		From:  from,
		To:    to,
		Value: value,
	})

	return nil
}

// PackageBlock package current transactions into a block
func (n *Node) PackageBlock(name string, metadata Metadata) Block {
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
