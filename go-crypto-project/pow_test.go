package main

import (
	"testing"
	"time"
)

func TestProofOfWork(t *testing.T) {
	// Test with a low difficulty
	prefix := "test"
	difficulty := 8
	
	start := time.Now()
	nonce, hash := Mine(prefix, difficulty)
	elapsed := time.Since(start)
	
	if nonce == 0 && hash == nil {
		t.Error("Failed to find a valid nonce for low difficulty")
	}
	
	t.Logf("Found nonce %d in %v", nonce, elapsed)
	
	// Verify the hash has the required number of leading zeros
	if !checkLeadingZeros(hash, difficulty) {
		t.Error("Hash does not have the required number of leading zeros")
	}
}

func TestProofOfWorkHigherDifficulty(t *testing.T) {
	// Test with a higher difficulty
	prefix := "harder"
	difficulty := 16
	
	start := time.Now()
	nonce, hash := Mine(prefix, difficulty)
	elapsed := time.Since(start)
	
	if nonce == 0 && hash == nil {
		t.Error("Failed to find a valid nonce for higher difficulty")
	}
	
	t.Logf("Found nonce %d in %v", nonce, elapsed)
	
	// Verify the hash has the required number of leading zeros
	if !checkLeadingZeros(hash, difficulty) {
		t.Error("Hash does not have the required number of leading zeros")
	}
}

func TestProofOfWorkVeryLowDifficulty(t *testing.T) {
	// Test with very low difficulty
	prefix := "easy"
	difficulty := 1
	
	start := time.Now()
	nonce, hash := Mine(prefix, difficulty)
	elapsed := time.Since(start)
	
	if nonce == 0 && hash == nil {
		t.Error("Failed to find a valid nonce for very low difficulty")
	}
	
	t.Logf("Found nonce %d in %v", nonce, elapsed)
	
	// Verify the hash has the required number of leading zeros
	if !checkLeadingZeros(hash, difficulty) {
		t.Error("Hash does not have the required number of leading zeros")
	}
}

func TestProofOfWorkZeroDifficulty(t *testing.T) {
	// Test with zero difficulty (should always succeed immediately)
	prefix := "zero"
	difficulty := 0
	
	start := time.Now()
	nonce, hash := Mine(prefix, difficulty)
	elapsed := time.Since(start)
	
	if nonce == 0 && hash == nil {
		t.Error("Failed to find a valid nonce for zero difficulty")
	}
	
	t.Logf("Found nonce %d in %v", nonce, elapsed)
	
	// Verify the hash has the required number of leading zeros
	if !checkLeadingZeros(hash, difficulty) {
		t.Error("Hash does not have the required number of leading zeros")
	}
}

func TestProofOfWorkDifferentPrefixes(t *testing.T) {
	// Test with different prefixes
	prefixes := []string{"hello", "world", "crypto", "lamport", "signature"}
	difficulty := 10
	
	for _, prefix := range prefixes {
		nonce, hash := Mine(prefix, difficulty)
		if nonce == 0 && hash == nil {
			t.Errorf("Failed to find a valid nonce for prefix '%s'", prefix)
			continue
		}
		
		if !checkLeadingZeros(hash, difficulty) {
			t.Errorf("Hash for prefix '%s' does not have the required number of leading zeros", prefix)
		}
		
		t.Logf("✓ Prefix '%s': nonce=%d", prefix, nonce)
	}
}

func TestProofOfWorkEdgeCases(t *testing.T) {
	// Test with edge case difficulties
	testCases := []struct {
		difficulty int
		description string
	}{
		{0, "zero difficulty"},
		{1, "single bit difficulty"},
		{8, "full byte difficulty"},
		{16, "two byte difficulty"},
	}
	
	for _, tc := range testCases {
		prefix := "edge"
		nonce, hash := Mine(prefix, tc.difficulty)
		
		if nonce == 0 && hash == nil {
			t.Errorf("Failed to find a valid nonce for %s", tc.description)
			continue
		}
		
		if !checkLeadingZeros(hash, tc.difficulty) {
			t.Errorf("Hash for %s does not have the required number of leading zeros", tc.description)
		}
		
		t.Logf("✓ %s: nonce=%d", tc.description, nonce)
	}
}

func BenchmarkProofOfWork(b *testing.B) {
	prefix := "benchmark"
	difficulty := 12
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mine(prefix, difficulty)
	}
}

func BenchmarkProofOfWorkLowDifficulty(b *testing.B) {
	prefix := "benchmark"
	difficulty := 8
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mine(prefix, difficulty)
	}
}

func BenchmarkProofOfWorkHighDifficulty(b *testing.B) {
	prefix := "benchmark"
	difficulty := 16
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mine(prefix, difficulty)
	}
}