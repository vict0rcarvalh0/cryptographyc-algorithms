package substitution

func SubstitutionEncrypt(text string) string {
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	another_alphabet := [26]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
	newPhrase := ""

	for _, char := range text {
		for i, v := range alphabet {
			if v == string(char) {
				newLetter := another_alphabet[i]
				newPhrase += newLetter
			}
		}
	}

	return newPhrase
}

func SubstitutionDecrypt(text string) string {
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	another_alphabet := [26]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
	newPhrase := ""

	for _, char := range text {
		for i, v := range another_alphabet {
			if v == string(char) {
				newLetter := alphabet[i]
				newPhrase += newLetter
			}
		}
	}

	return newPhrase
}
