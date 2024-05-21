package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, blockchain")

	// first, you need to create a new chain or
	// load an existing chain from the disk or network
	chain := NewChain(1)

	printJson("loaded chain", chain)

	// then, you start a node with the chain
	node := Node{
		Chain: chain,
	}

	printJson("started node", node)

	// now, you can execute transactions in a new block
	blocks := make(chan Block)
	stop := make(chan struct{})
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	// each block contains a list of transactions and metadata
	go func() {
		blocks <- node.NewBlock("Block 1", []Transaction{}, Metadata{})
		blocks <- node.NewBlock("Block 2", []Transaction{}, Metadata{})
		blocks <- node.NewBlock("Block 3", []Transaction{
			{
				From:  Wallet{Address: "me", Balance: Money{Amount: 100}},
				To:    Wallet{Address: "you", Balance: Money{Amount: 0}},
				Value: Money{Amount: 100},
			},
		}, Metadata{})
		stop <- struct{}{}
	}()

	// when a block is created, the node will try to "mine" it -
	// find a nonce that satisfies the difficulty level
	// then broadcast the block to the network to be added to the chain globally
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case block := <-blocks:
				printJson("created block", block)

				signed, _ := node.Nonce(block)

				printJson("signed block", signed)

				_ = node.Broadcast(signed)
				node.Chain.AddBlock(signed)

				fmt.Printf("block %s added to the chain\n", signed.Metadata["name"])
			case <-stop:
				return
			}
		}
	}()

	waitGroup.Wait()

	printTable(node.Chain)
}

func printJson(msg string, obj interface{}) {
	json, _ := json.Marshal(obj)
	fmt.Printf("%s: %v\n", msg, string(json))
}

func printTable(chain Chain) {
	for _, block := range chain.Blocks {
		fmt.Printf("%s\t%s\t%s\n", block.Metadata["name"], block.Hash, block.Timestamp)
	}
}
