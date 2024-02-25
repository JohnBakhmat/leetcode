package main

func isValid(s string) bool {

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []rune{}

	for _, r := range s {
		// If is open bracket put it on stack
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			// if closing bracket and there is no opening on stack
			return false
		} else {
			// if everythings nice pop the stack.
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
