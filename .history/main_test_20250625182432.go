package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

// cloneBlock returns a deep copy of the given block for test mutation
func cloneBlock(b Block) Block {
	return Block{
		Index:     b.Index,
		Timestamp: b.Timestamp,
		Data:      append([]byte{}, b.Data...),
		PrevHash:  append([]byte{}, b.PrevHash...),
		Nonce:     b.Nonce,
	}
}

// TestCalculateHash_AdversarialCollisions checks if different block contents produce unique hashes
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
		{
			"Delimiter Injection: Data contains PrevHash as prefix",
			base,
			func() Block {
				blk := cloneBlock(base)
				blk.Data = append(base.PrevHash, base.Data...)
				return blk
			}(),
		},
		{
			"Delimiter Injection: Data and PrevHash with null bytes",
			base,
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("foo\x00bar")
				blk.PrevHash = []byte("baz\x00qux")
				return blk
			}(),
		},
		{
			"Length Prefix Edge: Data and PrevHash same bytes, different length prefix",
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("AA")
				blk.PrevHash = []byte("A")
				return blk
			}(),
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("A")
				blk.PrevHash = []byte("AA")
				return blk
			}(),
		},
		{
			"Leading Zeros in Data: Data with and without leading zeros",
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("\x00\x00foobar")
				return blk
			}(),
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("foobar")
				return blk
			}(),
		},
		{
			"Unicode vs. ASCII: visually similar but different bytes",
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("é") // U+0065 U+0301
				return blk
			}(),
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("\u00e9") // U+00E9
				return blk
			}(),
		},
		{
			"All Fields Zero vs. Explicit Empty",
			func() Block {
				return Block{}
			}(),
			func() Block {
				return Block{
					Index:                 0,
					Timestamp:             0,
					Data:                  []byte{},
					PrevHash:              []byte{},
					Nonce:                 0,
					explicitlyInitialized: true,
				}
			}(),
		},
		{
			"Identical After Stripping Non-printables",
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("foo\nbar")
				return blk
			}(),
			func() Block {
				blk := cloneBlock(base)
				blk.Data = []byte("foo\rbar")
				return blk
			}(),
		},
		{
			"Data and PrevHash swapped, same combined bytes",
			func() Block {
				blk := cloneBlock(base)
				blk.Data, blk.PrevHash = base.PrevHash, base.Data
				return blk
			}(),
			base,
		},
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

// makeBlockchain creates a sample blockchain of the given size
func makeBlockchain(size int) []*Block {
	genesis := &Block{
		Index:     0,
		Timestamp: 0,
		Data:      []byte("Genesis"),
		PrevHash:  []byte{},
	}
	genesis.Hash = calculateHash(genesis)

	chain := []*Block{genesis}
	for i := 1; i < size; i++ {
		block := &Block{
			Index:     i,
			Timestamp: int64(i),
			Data:      []byte(fmt.Sprintf("Block %d", i)),
			PrevHash:  chain[i-1].Hash,
		}
		block.Hash = calculateHash(block)
		chain = append(chain, block)
	}
	return chain
}

// BenchmarkChainValidation measures performance of isChainValidCached
func BenchmarkChainValidation(b *testing.B) {
	sizes := []int{100, 1000, 5000, 10000}

	for _, size := range sizes {
		chain := makeBlockchain(size)
		b.Run(fmt.Sprintf("N=%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ok := isChainValidCached(chain)
				if !ok {
					b.Fatal("Chain is invalid, check your logic!")
				}
			}
		})
	}
}
