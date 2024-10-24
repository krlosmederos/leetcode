package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for i := 0; i < len(tokens); i++ {
		if op, isOp := ops[tokens[i]]; isOp {
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op(op2, op1))
		} else {
			toIns, _ := strconv.Atoi(tokens[i])
			stack = append(stack, toIns)
		}
	}

	return stack[0]
}
