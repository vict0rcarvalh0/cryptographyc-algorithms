package main

import (
	"fmt"
	"polialphabetic/vegenere"
	"time"
)

func main() {
	// Vigenere test
	plaintext := "Vigenere Cipher Performance Test"
	key := "KEY"

	startEncrypt := time.Now()
	encrypted := vegenere.EncryptVigenere(plaintext, key)
	durationEncrypt := time.Since(startEncrypt)

	startDecrypt := time.Now()
	decrypted := vegenere.DecryptVigenere(encrypted, key)
	durationDecrypt := time.Since(startDecrypt)

	fmt.Println("Texto Original:", plaintext)
	fmt.Println("Texto Cifrado:", encrypted)
	fmt.Println("Texto Decifrado:", decrypted)
	fmt.Printf("Tempo para Criptografar: %v\n", durationEncrypt)
	fmt.Printf("Tempo para Descriptografar: %v\n", durationDecrypt)
}
