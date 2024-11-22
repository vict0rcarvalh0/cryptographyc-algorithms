package vegenere

import "strings"

func EncryptVigenere(plaintext, key string) string {
	key = strings.ToLower(key)
	var result strings.Builder
	keyIndex := 0
	keyLength := len(key)

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			shift := int(key[keyIndex] - 'a')
			newChar := 'a' + (char-'a'+rune(shift))%26
			result.WriteRune(newChar)
			keyIndex = (keyIndex + 1) % keyLength
		} else if char >= 'A' && char <= 'Z' {
			shift := int(key[keyIndex] - 'a')
			newChar := 'A' + (char-'A'+rune(shift))%26
			result.WriteRune(newChar)
			keyIndex = (keyIndex + 1) % keyLength
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func DecryptVigenere(ciphertext, key string) string {
	key = strings.ToLower(key)
	var result strings.Builder
	keyIndex := 0
	keyLength := len(key)

	for _, char := range ciphertext {
		if char >= 'a' && char <= 'z' {
			shift := int(key[keyIndex] - 'a')
			newChar := 'a' + (char-'a'-rune(shift)+26)%26
			result.WriteRune(newChar)
			keyIndex = (keyIndex + 1) % keyLength
		} else if char >= 'A' && char <= 'Z' {
			shift := int(key[keyIndex] - 'a')
			newChar := 'A' + (char-'A'-rune(shift)+26)%26
			result.WriteRune(newChar)
			keyIndex = (keyIndex + 1) % keyLength
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
