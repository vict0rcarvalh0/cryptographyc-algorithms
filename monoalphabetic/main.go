package main

import (
	"fmt"
	"monoalphabetic/caesar"
	"monoalphabetic/substitution"
)

func main() {
	// Caesar
	text := "teste"
	shift := 3

	caesar_encryption := caesar.CeasarCypherEncript(text, shift)
	fmt.Println("Texto Cifrado:", caesar_encryption)

	caesar_decryption := caesar.CeasarDecypherEncript(caesar_encryption, shift)
	fmt.Println("Texto Defifrado:", caesar_decryption)

	caesar.VictorEncryptCaesar(text, shift)
	caesar.VictorDecryptCaesar(text, shift)

	// Substitution
	substitution_encryption := substitution.SubstitutionEncrypt(text)
	fmt.Println("Texto Cifrado:", substitution_encryption)

	substitution_decryption := substitution.SubstitutionDecrypt(text)
	fmt.Println("Texto Decifrado:", substitution_decryption)
}
