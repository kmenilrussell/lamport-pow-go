package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// PrivateKey represents a Lamport private key
type PrivateKey struct {
	Pairs [256][2][]byte
}

// PublicKey represents a Lamport public key
type PublicKey struct {
	Hashes [256][2][]byte
}

// GenerateKey generates a new Lamport key pair
func GenerateKey() (*PrivateKey, *PublicKey) {
	privateKey := &PrivateKey{}
	publicKey := &PublicKey{}
	
	for i := 0; i < 256; i++ {
		for j := 0; j < 2; j++ {
			// Generate a random 32-byte value for each private key element
			privKeyBytes := make([]byte, 32)
			if _, err := rand.Read(privKeyBytes); err != nil {
				panic(err) // In a real implementation, handle this error properly
			}
			privateKey.Pairs[i][j] = privKeyBytes
			
			// Hash the private key element to get the public key element
			hash := sha256.Sum256(privKeyBytes)
			publicKey.Hashes[i][j] = hash[:]
		}
	}
	
	return privateKey, publicKey
}

// Sign signs a message using the Lamport signature scheme
func Sign(message string, privateKey *PrivateKey) [256][]byte {
	// Hash the message to get a 256-bit digest
	messageHash := sha256.Sum256([]byte(message))
	var signature [256][]byte
	
	for i := 0; i < 256; i++ {
		// Determine which bit we're looking at
		byteIndex := i / 8
		bitIndex := uint(i % 8)
		bit := (messageHash[byteIndex] >> bitIndex) & 1
		
		// Select the appropriate half of the private key pair
		signature[i] = privateKey.Pairs[i][bit]
	}
	
	return signature
}

// Verify verifies a Lamport signature
func Verify(message string, signature [256][]byte, publicKey *PublicKey) bool {
	// Hash the message to get a 256-bit digest
	messageHash := sha256.Sum256([]byte(message))
	
	for i := 0; i < 256; i++ {
		// Determine which bit we're looking at
		byteIndex := i / 8
		bitIndex := uint(i % 8)
		bit := (messageHash[byteIndex] >> bitIndex) & 1
		
		// Hash the signature element
		sigHash := sha256.Sum256(signature[i])
		
		// Compare with the public key
		if !equalBytes(sigHash[:], publicKey.Hashes[i][bit]) {
			return false
		}
	}
	
	return true
}

// equalBytes compares two byte slices for equality
func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// StringToBytes converts a hex string to bytes
func StringToBytes(s string) []byte {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bytes
}

// ForgeSignature creates a forged signature using multiple signatures
func ForgeSignature(message string, messages []string, signatures [][256][]byte) [256][]byte {
	// Hash the message to get a 256-bit digest
	messageHash := sha256.Sum256([]byte(message))
	var forgedSignature [256][]byte
	
	// For each bit position, find a signature that reveals the required half
	for i := 0; i < 256; i++ {
		// Determine which bit we're looking at
		byteIndex := i / 8
		bitIndex := uint(i % 8)
		bit := (messageHash[byteIndex] >> bitIndex) & 1
		
		// Find a signature that reveals the required half
		for j := range signatures {
			// Hash the j-th message to get its digest
			sigMessageHash := sha256.Sum256([]byte(messages[j]))
			sigBit := (sigMessageHash[byteIndex] >> bitIndex) & 1
			
			// If this signature reveals the half we need, use it
			if sigBit == bit {
				forgedSignature[i] = signatures[j][i]
				break
			}
		}
	}
	
	return forgedSignature
}