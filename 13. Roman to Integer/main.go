package main

func romanToInt(s string) int {
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	for i := 0; i < len(s); i++ {
		c := roman[rune(s[i])]
		if i < len(s)-1 && c < roman[rune(s[i+1])] {
			total -= c
		} else {
			total += c
		}
	}
	return total
}
