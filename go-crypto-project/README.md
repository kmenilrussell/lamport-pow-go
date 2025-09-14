# Lamport Signatures and Proof-of-Work Implementation

This project implements a complete Lamport signature scheme with proof-of-work functionality in Go. It demonstrates the concepts of one-way functions, digital signatures, and computational work requirements.

## Features

### Part 1: Lamport Signatures
- **Key Generation**: Generate Lamport key pairs (private & public keys)
- **Signing**: Sign messages using the Lamport signature scheme
- **Verification**: Verify signatures against public keys
- **Hash-based**: Uses SHA-256 as the one-way hash function

### Part 2: Forgery Demonstration
- **Multi-signature vulnerability**: Shows how reusing private keys allows forgery
- **Signature construction**: Combines revealed values from multiple signatures
- **Automated testing**: Includes tests that verify forgery capabilities

### Part 3: Proof-of-Work
- **Configurable difficulty**: Adjustable number of leading zero bits required
- **Mining function**: Finds nonces that satisfy the difficulty requirement
- **Progress reporting**: Shows mining progress during computation
- **Integration**: Demonstrates POW alongside signature operations

## Files

- `main.go`: Main program with CLI interface and demonstrations
- `lamport.go`: Core Lamport signature implementation
- `pow.go`: Proof-of-work functionality
- `lamport_test.go`: Basic signature tests
- `forge_test.go`: Forgery scenario tests
- `pow_test.go`: Proof-of-work tests and benchmarks

## Usage

### Basic Signature Demonstration
```bash
go run main.go -difficulty 16
```

### Forgery Demonstration
```bash
go run main.go -forge -difficulty 16
```

### Running Tests
```bash
# Run all tests
go test

# Run specific test files
go test -run TestLamportSignature
go test -run TestForgery
go test -run TestProofOfWork

# Run with longer timeout for high-difficulty POW tests
go test -timeout 30m
```

### Running Benchmarks
```bash
go test -bench=.
go test -bench=BenchmarkProofOfWork
```

## Command Line Options

- `-difficulty N`: Set proof-of-work difficulty (number of leading zero bits, default: 16)
- `-forge`: Run forgery demonstration instead of basic signature demo

## Implementation Details

### Lamport Signatures
The Lamport signature scheme works as follows:

1. **Key Generation**: 
   - Private key: 256 pairs of random 32-byte values
   - Public key: SHA-256 hashes of all private key values

2. **Signing**:
   - Hash the message to get a 256-bit digest
   - For each bit in the digest, reveal the corresponding private key value
   - The signature consists of 256 revealed values

3. **Verification**:
   - Hash the message to get the same 256-bit digest
   - For each bit, hash the corresponding signature value
   - Compare with the public key to verify authenticity

### Forgery Scenario
When multiple signatures are created with the same private key:
- Each signature reveals half of each private key pair
- By combining signatures, an attacker can construct signatures for new messages
- The implementation demonstrates this by creating a forged message containing "forge" and an email address

### Proof-of-Work
The proof-of-work component:
- Takes a prefix and difficulty parameter D
- Finds a nonce such that `Hash(prefix || nonce)` has at least D leading zero bits
- Demonstrates computational effort required for different difficulty levels
- Can be integrated with signature operations to add work requirements

## Security Considerations

- **Key Reuse**: Lamport keys should only be used once for security
- **One-time Signatures**: This is a one-time signature scheme
- **Quantum Resistance**: Lamport signatures are quantum-resistant
- **Key Size**: Large key sizes (32KB for private key, 32KB for public key)

## Testing

The implementation includes comprehensive tests:

- **Basic functionality**: Key generation, signing, verification
- **Edge cases**: Different message types, key uniqueness
- **Forgery scenarios**: Various forgery attempts with different signature counts
- **Proof-of-work**: Different difficulty levels, edge cases, benchmarks
- **Performance**: Benchmarking for different difficulty settings

## Requirements

- Go 1.21 or later
- No external dependencies (uses only standard library)

## License

This project is for educational purposes.