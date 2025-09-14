package main

import (
	"testing"
)

func TestLamportSignature(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign a message
	message := "Test message"
	signature := Sign(message, privateKey)
	
	// Verify the signature
	if !Verify(message, signature, publicKey) {
		t.Error("Signature verification failed")
	}
	
	// Try to verify with a different message
	if Verify("Different message", signature, publicKey) {
		t.Error("Signature verification should have failed for a different message")
	}
}

func TestMultipleSignatures(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign multiple messages
	messages := []string{"Message 1", "Message 2", "Message 3"}
	signatures := make([][256][]byte, len(messages))
	
	for i, msg := range messages {
		signatures[i] = Sign(msg, privateKey)
	}
	
	// Verify all signatures
	for i, msg := range messages {
		if !Verify(msg, signatures[i], publicKey) {
			t.Errorf("Signature verification failed for message %d", i)
		}
	}
}

func TestKeyGeneration(t *testing.T) {
	// Generate multiple key pairs and ensure they're different
	privateKey1, publicKey1 := GenerateKey()
	privateKey2, publicKey2 := GenerateKey()
	
	// Private keys should be different
	samePrivate := true
	for i := 0; i < 256; i++ {
		for j := 0; j < 2; j++ {
			if !equalBytes(privateKey1.Pairs[i][j], privateKey2.Pairs[i][j]) {
				samePrivate = false
				break
			}
		}
		if !samePrivate {
			break
		}
	}
	
	if samePrivate {
		t.Error("Generated private keys should be different")
	}
	
	// Public keys should be different
	samePublic := true
	for i := 0; i < 256; i++ {
		for j := 0; j < 2; j++ {
			if !equalBytes(publicKey1.Hashes[i][j], publicKey2.Hashes[i][j]) {
				samePublic = false
				break
			}
		}
		if !samePublic {
			break
		}
	}
	
	if samePublic {
		t.Error("Generated public keys should be different")
	}
}

func TestSignatureUniqueness(t *testing.T) {
	// Generate a key pair
	privateKey, publicKey := GenerateKey()
	
	// Sign the same message twice
	message := "Same message"
	signature1 := Sign(message, privateKey)
	signature2 := Sign(message, privateKey)
	
	// Signatures should be identical for the same message
	for i := 0; i < 256; i++ {
		if !equalBytes(signature1[i], signature2[i]) {
			t.Error("Signatures for the same message should be identical")
			break
		}
	}
	
	// Verify both signatures
	if !Verify(message, signature1, publicKey) {
		t.Error("First signature verification failed")
	}
	if !Verify(message, signature2, publicKey) {
		t.Error("Second signature verification failed")
	}
}