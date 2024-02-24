package main

func singleNumber(nums []int) int {
	freq := make(map[int]int)

	for _, n := range nums {
		freq[n]++
	}

	for i, n := range freq {
		if n == 1 {
			return i
		}
	}
	return 0
}

func xor(nums []int) int {
	acc := 0
	for _, n := range nums {
		acc ^= n
	}
	return acc
}
