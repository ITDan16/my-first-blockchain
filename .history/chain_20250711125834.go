package main

import (
    "fmt"
    "time"
)

type Blockchain struct {
    blocks    []*Block
    config    Config
}

// Инициализация блокчейна с генезисом и конфигом
func NewBlockchain(config Config) (*Blockchain, error) {
    genesisBlock := &Block{
        Index:     0,
        Timestamp: time.Now().Unix(),
        Data:      []byte(config.GenesisData),
        PrevHash:  []byte{},
    }
    genesisBlock.Hash = calculateHash(genesisBlock)
    return &Blockchain{
        blocks: []*Block{genesisBlock},
        config: config,
    }, nil
}

// Добавление блока с валидацией
func (bc *Blockchain) AddBlock(data string) error {
    prevBlock := bc.blocks[len(bc.blocks)-1]
    newBlock := &Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now().Unix(),
        Data:      []byte(data),
        PrevHash:  prevBlock.Hash,
    }
    hash, nonce := proofOfWork(newBlock, bc.config.Difficulty)
    newBlock.Hash = hash
    newBlock.Nonce = nonce

    // Проверка нового блока
    if err := ValidateBlock(prevBlock, newBlock, bc.config); err != nil {
        return fmt.Errorf("block validation failed: %w", err)
    }

    bc.blocks = append(bc.blocks, newBlock)
    return nil
}

// Получение блока по индексу
func (bc *Blockchain) GetBlock(index int) (*Block, error) {
    if index < 0 || index >= len(bc.blocks) {
        return nil, fmt.Errorf("block with index %d not found", index)
    }
    return bc.blocks[index], nil
}

// Проверка всей цепочки
func (bc *Blockchain) Validate() error {
    for i := 1; i < len(bc.blocks); i++ {
        prev := bc.blocks[i-1]
        curr := bc.blocks[i]
        if err := ValidateBlock(prev, curr, bc.config); err != nil {
            return fmt.Errorf("chain invalid at block %d: %w", curr.Index, err)
        }
    }
    return nil
}