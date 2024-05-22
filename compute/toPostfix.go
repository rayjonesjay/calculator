package compute

import (
	"calculator/parse"
	"fmt"
)

// InfixToPostFix() converts from infix notation a + b  -> postfix notation a b +
func InfixToPostFix(tokens []string) ([]string, error) {
	var postfix []string
	var stack []string

	for _, token := range tokens {
		if parse.IsNumber(token) {
			postfix = append(postfix, token)

		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // Pop "("
		} else { // Operator
			for len(stack) > 0 && Precedence(stack[len(stack)-1]) >= Precedence(token) {

				postfix = append(postfix, stack[len(stack)-1])
	
				stack = stack[:len(stack)-1]

			}

			//other oprators like - + * 
			stack = append(stack, token)

		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])

		stack = stack[:len(stack)-1]
	}

	return postfix, nil
}
