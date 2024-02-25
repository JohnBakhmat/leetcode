package main

func longestCommonPrefix(strs []string) string {
	dump := make([]rune, 0, 200)
	minLength := 200
	for _, word := range strs {
		if minLength > len(word) {
			minLength = len(word)
		}
	}

	for i := 0; i < minLength; i++ {
		cur := strs[0][i]
		for _, word := range strs {
			if cur != word[i] {
				return string(dump)
			}
		}
		dump = append(dump, rune(cur))
	}

	return string(dump)
}
