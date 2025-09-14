# 🔐 Lamport Signatures & Proof-of-Work Implementation

A modern Go implementation demonstrating **quantum-resistant cryptography** through Lamport signatures combined with proof-of-work security mechanisms.

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Quantum Resistant](https://img.shields.io/badge/Quantum-Resistant-orange)

## ✨ Features

### 🔏 Lamport Signatures
- **Key Generation**: Create cryptographically secure Lamport key pairs
- **Digital Signatures**: Sign messages using one-time signature scheme
- **Signature Verification**: Validate message authenticity with public keys
- **SHA-256 Based**: Built on Go's standard cryptographic hash function

### ⛏️ Proof-of-Work System
- **Adjustable Difficulty**: Configure leading zero bits requirement
- **Real-time Mining**: Visual progress reporting during computation
- **Nonce Discovery**: Find valid proofs that meet difficulty targets
- **Seamless Integration**: Combined signature and PoW workflow

### ⚠️ Security Demonstrations
- **Key Reuse Vulnerability**: Shows practical forgery attacks
- **Signature Construction**: How attackers combine revealed private values
- **Educational Tests**: Validate security assumptions and limitations

## 📁 Project Structure

| File | Purpose |
|------|---------|
| `main.go` | CLI interface and demonstration flows |
| `lamport.go` | Core Lamport signature implementation |
| `pow.go` | Proof-of-work mining functionality |
| `lamport_test.go` | Unit tests for signature operations |
| `forge_test.go` | Signature forgery demonstration tests |
| `pow_test.go` | PoW algorithm tests and benchmarks |

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or newer
- No external dependencies required

### Basic Signature Demonstration
```bash
go run main.go -difficulty 16
```

### Forgery Attack Demonstration
```bash
go run main.go -forge -difficulty 16
```

### Testing
```bash
# Run all tests
go test

# Run specific test suites
go test -run TestLamportSignature
go test -run TestForgery
go test -run TestProofOfWork

# Extended timeout for high-difficulty tests
go test -timeout 30m
```

### Benchmarking
```bash
# Run all benchmarks
go test -bench=.

# Specific PoW benchmarking
go test -bench=BenchmarkProofOfWork
```

## ⚙️ Configuration

| Flag | Default | Description |
|------|---------|-------------|
| `-difficulty` | 16 | Number of leading zero bits required for PoW |
| `-forge` | false | Run forgery demonstration instead of basic flow |

## 🔍 Technical Implementation

### Lamport Signature Scheme
- **Private Key**: 256 pairs of random 32-byte values
- **Public Key**: SHA-256 hashes of all private values
- **Signing**: Reveals private values based on message hash bits
- **Verification**: Hashes signature values and compares with public key

### Proof-of-Work Algorithm
- **Input**: Message prefix and difficulty target
- **Process**: Iterative nonce discovery through brute force
- **Output**: Valid nonce producing hash with required leading zeros
- **Difficulty**: Exponential computation time increase per bit

### Forgery Vulnerability
- **Cause**: Private key reuse across multiple signatures
- **Method**: Combining revealed values from different signatures
- **Result**: Ability to forge signatures for new messages
- **Prevention**: Strict one-time usage of private keys

## 🛡️ Security Considerations

> **Critical Warning**: Lamport signatures are **one-time use only**

- **Key Reuse**: Leads to immediate vulnerability to forgery attacks
- **Quantum Resistance**: Secure against Shor's and Grover's algorithms
- **Key Size**: Large keys (~16KB) and signatures (~32KB) required
- **Practical Limitations**: Generally unsuitable for general-purpose use

## 📚 Educational Value

This implementation serves as an excellent learning resource for:
- Post-quantum cryptography concepts
- One-way function properties and applications
- Digital signature construction and verification
- Proof-of-work consensus mechanisms
- Cryptographic security vulnerability analysis

## 🧪 Testing Philosophy

The test suite validates:
- Basic signature functionality under normal conditions
- Boundary cases and error conditions
- Proof-of-work difficulty compliance
- Intentional vulnerability demonstrations
- Performance characteristics and benchmarks

---

**Note**: This implementation is for educational purposes and should not be used in production systems without extensive security review and modifications.
