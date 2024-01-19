package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if token == "+" {
				stack = append(stack, num1+num2)
			} else if token == "-" {
				stack = append(stack, num1-num2)
			} else if token == "*" {
				stack = append(stack, num1*num2)
			} else if token == "/" {
				stack = append(stack, num1/num2)
			}
		} else {
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
