````markdown
# Lamport Signatures & Proof-of-Work Implementation

This project implements a complete **Lamport signature scheme** with **proof-of-work** (PoW) functionality in Go. It's designed to demonstrate the fundamental concepts of one-way functions, digital signatures, and computational work requirements in a practical context.

---

## Features

### Lamport Signatures
- **Key Generation**: Generate Lamport key pairs (private and public).
- **Signing & Verification**: Sign messages and verify them using the Lamport one-time signature scheme.
- **Hash-based**: Leverages the **SHA-256** hash function as its core one-way function.

### Proof-of-Work
- **Configurable Difficulty**: Adjust the number of leading zero bits required for a valid hash.
- **Mining Function**: Finds the correct **nonce** that satisfies the difficulty requirement.
- **Progress Reporting**: Displays mining progress during computation for a better user experience.
- **Integration**: Demonstrates how PoW can be combined with signature operations.

### Forgery Demonstration
- **Multi-signature Vulnerability**: Explicitly shows how reusing a private key allows for signature forgery.
- **Signature Construction**: Combines revealed values from multiple signatures to create a new, forged signature.
- **Automated Testing**: Includes tests that validate the forgery capabilities, highlighting a crucial security flaw.

---

## Files

| File | Purpose |
|---|---|
| `main.go` | The main program, providing a command-line interface and demonstration flows. |
| `lamport.go` | The core Lamport signature implementation, including key generation, signing, and verification. |
| `pow.go` | The proof-of-work functionality, including mining and difficulty handling. |
| `lamport_test.go` | Unit tests for the basic signature functionality. |
| `forge_test.go` | Tests for the forgery scenario, exposing the misuse vulnerability. |
| `pow_test.go` | Unit tests and benchmarks for the proof-of-work component. |

---

## Usage

### Run a Basic Signature Demonstration
```bash
go run main.go -difficulty 16
````

### Run the Forgery Demonstration

```bash
go run main.go -forge -difficulty 16
```

### Running Tests

```bash
# Run all tests
go test

# Run specific tests
go test -run TestLamportSignature
go test -run TestForgery
go test -run TestProofOfWork

# Run with a longer timeout (for higher PoW difficulty)
go test -timeout 30m
```

### Running Benchmarks

```bash
go test -bench=.
go test -bench=BenchmarkProofOfWork
```

-----

## Command Line Options

| Flag | Description |
|---|---|
| `-difficulty N` | Sets the proof-of-work difficulty as the number of leading zero bits. The default is `16`. |
| `-forge` | Runs the forgery demonstration instead of the basic signature flow. |

-----

## Implementation Details

### Lamport Signatures

  - **Key Generation**: Generates 256 pairs of random 32-byte values for the private key. The public key is created by hashing each of these 512 values.
  - **Signing**:
    1.  Hash the message using SHA-256 to get a 256-bit digest.
    2.  For each bit of the digest, the corresponding value from the private key pair is "revealed" as the signature.
  - **Verification**:
    1.  Re-hash the message.
    2.  For each bit, hash the revealed signature value and compare it against the corresponding value in the public key.

### Forgery Scenario

  - When a Lamport key is used to sign multiple messages, parts of the private key are revealed with each new signature.
  - An attacker can combine revealed values from different signatures to construct a new, forged signature for a different message.
  - The demonstration ensures the forged message contains the word `"forge"` to make the attack explicit.

### Proof-of-Work

  - The mining function takes a **prefix** (e.g., derived from a message or public key) and a difficulty parameter `D`.
  - It searches for a **nonce** `n` such that the hash of the combined prefix and nonce (`SHA256(prefix || nonce)`) has at least `D` leading zero bits.
  - This process requires significant computational work, which increases exponentially with difficulty.

-----

## Security Considerations

  - **Key Reuse**: Lamport signatures are a **one-time signature scheme**. Reusing a private key is a critical security vulnerability that allows for forgery.
  - **Quantum Resistance**: Because it relies solely on one-way hash functions, the Lamport scheme is considered **quantum-resistant** to certain attacks, such as those by Shor's or Grover's algorithms.
  - **Key & Signature Size**: A major drawback is the large size of the keys (\~32 KB each for private and public) and signatures (\~16 KB), making it impractical for many real-world applications.

-----

## Requirements

  - Go version **1.21** or later
  - No external dependencies—the project uses only the Go standard library.

<!-- end list -->

```
```
