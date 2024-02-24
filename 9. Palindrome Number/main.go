package main

import (
	"strconv"
	"strings"
)

func easy(x int) bool {
	x_str := strings.Split(strconv.Itoa(x), "")
	n := len(x_str) - 1
	for i, l := range x_str {
		if l != x_str[n-i] {
			return false
		}
	}
	return true
}


// Chip away start and end of number and compare. 
func without_string(x int) bool {
	if (x > 0 && x%10 == 0) || x < 0 {
		return false
	}

	y := 0
	for y < x {
		y = y*10 + x%10
		x /= 10
	}

	return x == y || x == y/10
}
