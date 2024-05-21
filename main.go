package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, blockchain")

	// first, you need to create a new chain or
	// load an existing chain from the disk or network
	chain := NewChain(5)

	printJson("loaded chain", chain)

	// then, you start a node with the chain
	node := Node{
		Chain: chain,
	}

	printJson("started node", node)

	for i := 1; i < 4; i++ {
		// now, you can submit transactions that will be queued in a node
		node.SubmitTransaction(Wallet{Address: "me"}, Wallet{Address: "you"}, Money{Amount: 100})

		// worker node will pack transactions into a block
		block := node.PackageBlock(fmt.Sprintf("Block %d", i), Metadata{})

		printJson("created block", block)

		// then, the block will be signed and broadcasted to the network
		signed, _ := node.Nonce(block)

		printJson("signed block", signed)

		_ = node.Broadcast(signed)
		
	}

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
