package main

import (
	"strings"
	"testing"
)

// TestForgery tests the ability to forge a signature when multiple signatures are available
func TestForgery(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign multiple messages with the same private key
	messages := []string{"Message 1", "Message 2", "Message 3", "Message 4"}
	signatures := make([][256][]byte, len(messages))
	
	for i, msg := range messages {
		signatures[i] = Sign(msg, privateKey)
	}
	
	// Forge a signature for a new message
	forgedMessage := "I am forger@example.com forging this message"
	forgedSignature := ForgeSignature(forgedMessage, messages, signatures)
	
	// Verify the forged signature
	if !Verify(forgedMessage, forgedSignature, publicKey) {
		t.Error("Forged signature verification failed")
	}
	
	// Check that the forged message contains "forge" and an email
	if !strings.Contains(forgedMessage, "forge") {
		t.Error("Forged message should contain the word 'forge'")
	}
	if !strings.Contains(forgedMessage, "@") {
		t.Error("Forged message should contain an email address")
	}
}

// TestForgeryWithMinimalSignatures tests forging with the minimum number of signatures
func TestForgeryWithMinimalSignatures(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign only 2 messages (minimum for potential forgery)
	messages := []string{"Message A", "Message B"}
	signatures := make([][256][]byte, len(messages))
	
	for i, msg := range messages {
		signatures[i] = Sign(msg, privateKey)
	}
	
	// Attempt to forge a signature
	forgedMessage := "This is a forged message by test@example.com"
	forgedSignature := ForgeSignature(forgedMessage, messages, signatures)
	
	// Verify the forged signature (may or may not work depending on bit patterns)
	valid := Verify(forgedMessage, forgedSignature, publicKey)
	
	// The forgery might not always work with only 2 signatures, but when it does, it should meet requirements
	if valid {
		if !strings.Contains(forgedMessage, "forged") {
			t.Error("Forged message should contain the word 'forged'")
		}
		if !strings.Contains(forgedMessage, "@") {
			t.Error("Forged message should contain an email address")
		}
		t.Log("Forgery succeeded with 2 signatures")
	} else {
		t.Log("Forgery failed with 2 signatures (expected behavior)")
	}
}

// TestForgeryWithDifferentMessages tests forging with various message types
func TestForgeryWithDifferentMessages(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign messages with different content
	messages := []string{
		"The quick brown fox jumps over the lazy dog",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
		"1234567890!@#$%^&*()",
		"Cryptography is the practice of secure communication",
	}
	signatures := make([][256][]byte, len(messages))
	
	for i, msg := range messages {
		signatures[i] = Sign(msg, privateKey)
	}
	
	// Forge signatures for different types of messages
	testCases := []struct {
		message string
		description string
	}{
		{"I am hacker@evil.com attempting to forge this", "hacker claim"},
		{"This is a forged document by forger@fake.com", "document forgery"},
		{"Forged transaction from thief@scam.org", "transaction forgery"},
	}
	
	for _, tc := range testCases {
		forgedSignature := ForgeSignature(tc.message, messages, signatures)
		if Verify(tc.message, forgedSignature, publicKey) {
			t.Logf("✓ Successfully forged: %s", tc.description)
		} else {
			t.Logf("✗ Failed to forge: %s", tc.description)
		}
	}
}

// TestForgeryImpossibleWithSingleSignature tests that forgery is impossible with only one signature
func TestForgeryImpossibleWithSingleSignature(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign only one message
	messages := []string{"Single message"}
	signatures := make([][256][]byte, len(messages))
	signatures[0] = Sign(messages[0], privateKey)
	
	// Attempt to forge a signature for a different message
	forgedMessage := "This should not be forgeable with one signature"
	forgedSignature := ForgeSignature(forgedMessage, messages, signatures)
	
	// The forged signature should not verify correctly
	if Verify(forgedMessage, forgedSignature, publicKey) {
		t.Error("Forgery should not be possible with only one signature")
	}
	
	// But the original message should still verify
	if !Verify(messages[0], signatures[0], publicKey) {
		t.Error("Original signature should still verify correctly")
	}
}