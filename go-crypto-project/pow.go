package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

// Mine finds a nonce such that Hash(prefix || nonce) has at least D leading zero bits
func Mine(prefix string, difficulty int) (uint64, []byte) {
	var nonce uint64
	maxNonce := uint64(math.MaxUint64)
	reportInterval := uint64(1000000) // Report progress every million nonces
	nextReport := reportInterval
	
	// Convert prefix to bytes
	prefixBytes := []byte(prefix)
	
	// Precompute the target mask
	fullBytes := difficulty / 8
	remainingBits := difficulty % 8
	var targetMask byte
	if remainingBits > 0 {
		targetMask = byte(0xFF) << (8 - remainingBits)
	}
	
	for nonce = 0; nonce < maxNonce; nonce++ {
		// Report progress periodically
		if nonce == nextReport {
			fmt.Printf("Tried %d nonces...\n", nonce)
			nextReport += reportInterval
		}
		
		// Create a buffer for the nonce
		nonceBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(nonceBytes, nonce)
		
		// Concatenate prefix and nonce
		data := append(prefixBytes, nonceBytes...)
		
		// Hash the data
		hash := sha256.Sum256(data)
		
		// Check if the hash has at least D leading zero bits
		valid := true
		for i := 0; i < fullBytes; i++ {
			if hash[i] != 0 {
				valid = false
				break
			}
		}
		
		if valid && remainingBits > 0 {
			if (hash[fullBytes] & targetMask) != 0 {
				valid = false
			}
		}
		
		if valid {
			return nonce, hash[:]
		}
	}
	
	return 0, nil // No nonce found (unlikely for reasonable difficulty)
}

// PrintPOW prints the proof-of-work result
func PrintPOW(prefix string, difficulty int) {
	fmt.Printf("Mining with prefix '%s' and difficulty %d...\n", prefix, difficulty)
	nonce, hash := Mine(prefix, difficulty)
	if nonce == 0 && hash == nil {
		fmt.Println("Failed to find a valid nonce")
		return
	}
	fmt.Printf("Found nonce: %d\n", nonce)
	fmt.Printf("Hash: %x\n", hash)
}

// checkLeadingZeros checks if a hash has at least D leading zero bits
func checkLeadingZeros(hash []byte, difficulty int) bool {
	if difficulty <= 0 {
		return true
	}
	
	// Check full bytes first
	fullBytes := difficulty / 8
	for i := 0; i < fullBytes; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	
	// Check remaining bits
	remainingBits := difficulty % 8
	if remainingBits > 0 {
		mask := byte(0xFF) << (8 - remainingBits)
		if (hash[fullBytes] & mask) != 0 {
			return false
		}
	}
	
	return true
}