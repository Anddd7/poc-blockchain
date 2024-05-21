# poc-blockchain

## Core Model

Block

- hash
- prehash
- transactions

Chain

- blocks

Node

- chain

## Workflow

1. Start nodes and chain

- There are multiple nodes that run the same chain
- Chain is started with a genesis block

2. Submit transactions

- User submit a transaction to a node
- Node will validate the transaction
- Node will put the transaction in mempool
- Node will broadcast the transaction to other nodes and wait for package

3. Mine block

- Miner node will package transactions in mempool into a block
- Miner node will evaluate the PoW util finding a valid block
- Miner node will broadcast the block to other nodes
- Other node will validate the block and add it to the chain
