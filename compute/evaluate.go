package compute

import (
	"fmt"
	"math"
	"strconv"
	"calculator/parse"
)

func EvaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if parse.IsNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			case "^":
				result = math.Pow(a, b)
			default:
				return 0, fmt.Errorf("invalid operator: %s", token)
			}
			
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack[len(stack)-1], nil
}
