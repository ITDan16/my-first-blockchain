package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

// Клонирует блок (для мутаций в тестах)
func cloneBlock(b Block) Block {
	return Block{
		Index:     b.Index,
		Timestamp: b.Timestamp,
		Data:      append([]byte{}, b.Data...),
		PrevHash:  append([]byte{}, b.PrevHash...),
		Nonce:     b.Nonce,
	}
}

// Проверка коллизий хеша для разных блоков
func TestCalculateHash_AdversarialCollisions(t *testing.T) {
	base := Block{
		Index:     5,
		Timestamp: 1111222233,
		Data:      []byte("foo|bar||baz"),
		PrevHash:  []byte("feedcafe"),
		Nonce:     1337,
	}
	calculateHash(&base)

	cases := []struct {
		name string
		a, b Block
	}{
		// ... (оставь все твои кейсы без изменений)
		// они абсолютно совместимы
        // (можно просто скопировать все твои кейсы сюда)
		{
			"Delimiter Injection: Data contains PrevHash as prefix",
			base,
			func() Block {
				blk := cloneBlock(base)
				blk.Data = append(base.PrevHash, base.Data...)
				return blk
			}(),
		},
        // ... остальные кейсы без изменений
		{
			"Different Index but other fields match",
			func() Block {
				blk := cloneBlock(base)
				blk.Index++
				return blk
			}(),
			base,
		},
	}

	for _, tc := range cases {
		hashA := calculateHash(&tc.a)
		hashB := calculateHash(&tc.b)
		if bytes.Equal(hashA, hashB) {
			var buf bytes.Buffer
			buf.WriteString("Hash collision detected for case '" + tc.name + "':\n")
			buf.WriteString(fmt.Sprintf("Block A: %+v\n", tc.a))
			buf.WriteString(fmt.Sprintf("Block B: %+v\n", tc.b))
			buf.WriteString("Hash: " + hex.EncodeToString(hashA) + "\n")
			buf.WriteString("BlockA bytes: " + hex.EncodeToString(tc.a.Data) + "\n")
			buf.WriteString("BlockB bytes: " + hex.EncodeToString(tc.b.Data) + "\n")
			t.Error(buf.String())
		}
	}
}

// Создает новый блокчейн, используя новые структуры и методы
func makeBlockchain(size int, config Config) *Blockchain {
	bc, err := NewBlockchain(config)
	if err != nil {
		panic("failed to create blockchain: " + err.Error())
	}
	for i := 1; i < size; i++ {
		if err := bc.AddBlock(fmt.Sprintf("Block %d", i)); err != nil {
			panic("failed to add block: " + err.Error())
		}
	}
	return bc
}

// BenchmarkChainValidation измеряет производительность Validate
func BenchmarkChainValidation(b *testing.B) {
	sizes := []int{100, 1000, 5000, 10000}
	config := Config{
		GenesisData: "Genesis",
		Difficulty:  1,   // чтобы не тормозить PoW в бенчмарке
		HashLength:  32,
	}

	for _, size := range sizes {
		bc := makeBlockchain(size, config)
		b.Run(fmt.Sprintf("N=%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := bc.Validate(); err != nil {
					b.Fatal("Chain is invalid, check your logic! " + err.Error())
				}
			}
		})
	}
}