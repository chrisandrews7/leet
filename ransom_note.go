package main

func canConstruct(ransomNote string, magazine string) bool {
	lettersUsed := make(map[rune]int)

	for _, letter := range magazine {
		lettersUsed[letter] += 1
	}

	for _, letter := range ransomNote {
		if usesLeft, ok := lettersUsed[letter]; ok && usesLeft > 0 {
			lettersUsed[letter]--
			continue
		}

		return false
	}

	return true
}
