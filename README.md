# Minimal Go Proof-of-Work Blockchain

A minimalist, readable, and extensible Go blockchain with Proof-of-Work support, typed errors, tests, and clean architecture.

---

## 🚀 Features

- Fully modular architecture (main, block, chain, validation, config)
- Configurable Proof-of-Work difficulty
- Block generation and validation
- Clean, typed error handling
- Easily extendable to P2P or transactions
- Includes unit tests and benchmarks

---

## 📦 Quick Start

### Clone the repo:

```sh
git clone https://github.com/yourname/go-pow-blockchain.git
cd go-pow-blockchain
```

### Run the application:

```sh
go run main.go
```

**Sample output:**

```
Blockchain:
Index: 0, Data: Genesis, Hash: 3c96f6d9e1...
Index: 1, Data: Second Block, Hash: 0000fa2cf2...
Index: 2, Data: Third Block, Hash: 00003ebc45...

Blockchain is valid!
```

---

## ⚙️ Configuration

- All parameters are centralized in `config.go`:
  - `GenesisData` — genesis block data
  - `Difficulty` — PoW difficulty (number of leading zeros in hash)
  - `HashLength` — hash length (reserved for future extensions)

---

## 🛠️ Project Structure

```
main.go          // Entry point, demo usage
block.go         // Block structure, serialization, PoW, hashing
chain.go         // Blockchain structure and methods (add, validate)
validation.go    // Validation functions and typed errors
config.go        // Project configuration
main_test.go     // Tests and benchmarks
go.mod           // Go modules
```

---

## 📚 API Usage Example

```go
config := Config{
    GenesisData: "Genesis",
    Difficulty:  4,
    HashLength:  32,
}
bc, err := NewBlockchain(config)
if err != nil { log.Fatal(err) }

if err := bc.AddBlock("Some data"); err != nil {
    log.Fatal(err)
}
if err := bc.Validate(); err != nil {
    log.Fatalf("Invalid chain: %v", err)
}
```

---

## ✅ Tests & Benchmarks

Run tests and benchmarks:

```sh
go test -v -bench=.
```

- Unit tests check hash uniqueness, blockchain validity, and fault tolerance.
- Benchmarks measure chain validation speed at various lengths.

---

## 🧩 Ideas for Extension

- P2P networking with libp2p or net/http
- Transactions and wallet balances
- Web interface and REST API
- Proof-of-Stake algorithms
- Database integration

---

## 🤝 Contributing

Pull requests and ideas are welcome! Best practices:

- Keep code style and modularity
- Add tests for new features
- Describe changes in the changelog

---

## 📜 License

MIT

---

> **Author:** [Danylo Mozhaiev](https://github.com/yourname)
> *Inspired by Arec1b0, Go learning projects and blockchain principles.*
> *Built for learning and fun. Want to improve something? Just open an issue or PR!*


