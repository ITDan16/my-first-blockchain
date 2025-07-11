package main

import (
	"fmt"
	"time"
)

func main() {
	config := Config{
		GenesisData: "Genesis",
		Difficulty:  4,
	}

	genesisBlock := &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      []byte(config.GenesisData),
		PrevHash:  []byte{},
	}
	genesisBlock.Hash = calculateHash(genesisBlock)

	bc := &Blockchain{
		Blocks:     []*Block{genesisBlock},
		Difficulty: config.Difficulty,
	}

	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")

	fmt.Println("Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d, Data: %s, Hash: %x...\n", block.Index, string(block.Data), block.Hash[:10])
	}

	fmt.Printf("\nIs blockchain valid? %t\n", bc.IsValid())
}
