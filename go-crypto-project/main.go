package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Parse command line flags
	difficulty := flag.Int("difficulty", 16, "Proof-of-work difficulty (number of leading zero bits)")
	demoForge := flag.Bool("forge", false, "Run forgery demonstration")
	flag.Parse()
	
	// Validate difficulty
	if *difficulty < 0 || *difficulty > 256 {
		fmt.Println("Difficulty must be between 0 and 256")
		os.Exit(1)
	}
	
	if *demoForge {
		demonstrateForgery(*difficulty)
	} else {
		demonstrateBasicSignature(*difficulty)
	}
}

func demonstrateBasicSignature(difficulty int) {
	fmt.Println("=== Basic Lamport Signature Demonstration ===")
	
	// Generate a key pair
	fmt.Println("Generating key pair...")
	privateKey, publicKey := GenerateKey()
	
	// Sign a message
	message := "Hello, Lamport signatures!"
	fmt.Printf("Signing message: %s\n", message)
	signature := Sign(message, privateKey)
	
	// Verify the signature
	valid := Verify(message, signature, publicKey)
	fmt.Printf("Verify worked? %v\n", valid)
	
	// Demonstrate proof-of-work
	fmt.Println("\n=== Proof-of-Work Demonstration ===")
	prefix := fmt.Sprintf("%x", publicKey.Hashes[0][0]) // Use part of the public key as prefix
	PrintPOW(prefix, difficulty)
}

func demonstrateForgery(difficulty int) {
	fmt.Println("=== Forgery Demonstration ===")
	
	// Generate a key pair
	fmt.Println("Generating key pair...")
	privateKey, publicKey := GenerateKey()
	
	// Sign multiple messages with the same private key
	messages := []string{"Message 1", "Message 2", "Message 3", "Message 4"}
	signatures := make([][256][]byte, len(messages))
	
	fmt.Println("Signing multiple messages with the same private key...")
	for i, msg := range messages {
		fmt.Printf("Signing: %s\n", msg)
		signatures[i] = Sign(msg, privateKey)
	}
	
	// Forge a signature for a new message
	forgedMessage := "I am forger@example.com forging this message"
	fmt.Printf("\nForging signature for message: %s\n", forgedMessage)
	
	forgedSignature := ForgeSignature(forgedMessage, messages, signatures)
	
	// Verify the forged signature
	fmt.Println("Verifying forged signature...")
	valid := Verify(forgedMessage, forgedSignature, publicKey)
	fmt.Printf("Forged signature verification: %v\n", valid)
	
	// Check requirements
	if strings.Contains(forgedMessage, "forge") && strings.Contains(forgedMessage, "@") {
		fmt.Println("✓ Forged message contains 'forge' and email address")
	} else {
		fmt.Println("✗ Forged message does not meet requirements")
	}
	
	// Demonstrate proof-of-work for the forged signature
	fmt.Println("\n=== Proof-of-Work for Forged Signature ===")
	prefix := fmt.Sprintf("%x", publicKey.Hashes[0][0]) // Use part of the public key as prefix
	PrintPOW(prefix, difficulty)
}