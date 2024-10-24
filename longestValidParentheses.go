package main

func longestValidParentheses(s string) int {
	ans := 0
	var stack []int

	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
				ans = max(ans, i-stack[len(stack)-1])
			} else {
				stack[0] = i
			}
		}
	}

	return ans
}
