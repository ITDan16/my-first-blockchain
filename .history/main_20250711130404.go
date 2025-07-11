package main

import (
    "fmt"
    "log"
)

func main() {
    config := Config{
        GenesisData: "Genesis",
        Difficulty:  4,
        HashLength:  32, 
    blockchain, err := NewBlockchain(config)
    if err != nil {
        log.Fatalf("Failed to initialize blockchain: %v", err)
    }

    if err := blockchain.AddBlock("Second Block"); err != nil {
        log.Fatalf("Failed to add block: %v", err)
    }
    if err := blockchain.AddBlock("Third Block"); err != nil {
        log.Fatalf("Failed to add block: %v", err)
    }

    fmt.Println("Blockchain:")
    for _, block := range blockchain.blocks {
        fmt.Printf("Index: %d, Data: %s, Hash: %s\n",
            block.Index, string(block.Data), fmt.Sprintf("%x", block.Hash)[:10]+"...")
    }

    fmt.Println()
    if err := blockchain.Validate(); err != nil {
        log.Printf("Blockchain INVALID: %v", err)
    } else {
        fmt.Println("Blockchain is valid!")
    }
}