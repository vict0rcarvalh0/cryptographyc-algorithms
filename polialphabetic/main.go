package main

import (
	"fmt"
	"polialphabetic/vegenere"
	"time"
)

func main() {
	plaintext := "teste"
	key := "qwertyuiop"

	startEncrypt := time.Now()
	encrypted := vegenere.EncryptVigenere(plaintext, key)
	durationEncrypt := time.Since(startEncrypt)
	fmt.Println("Texto Cifrado:", encrypted)

	startDecrypt := time.Now()
	decrypted := vegenere.DecryptVigenere(encrypted, key)
	durationDecrypt := time.Since(startDecrypt)
	fmt.Println("Texto Decifrado:", decrypted)

	fmt.Printf("Tempo para Criptografar: %v\n", durationEncrypt)
	fmt.Printf("Tempo para Descriptografar: %v\n", durationDecrypt)
}
