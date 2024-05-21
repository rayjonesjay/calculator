package compute

import ("calculator/parse")

// InfixToPostFix() converts from infix notation a + b  -> postfix notation a b +
func InfixToPostFix(tokens []string) ([]string, error) {
	var postfix []string

	var stack []string

	lengthOfStack := len(stack)

	for _, token := range tokens {

		if parse.IsNumber(token) {

			postfix = append(postfix, token)

		} else if token == "(" {

			stack = append(stack, token)

		} else if token == ")" {

			for len(stack) > 0 && stack[lengthOfStack-1] != "(" {

				// append the top element
				postfix = append(postfix, stack[lengthOfStack-1])

				// slice to remove the top element from stack
				stack = stack[:lengthOfStack-1]

			}

			stack = stack[:lengthOfStack-1] // remove the (
		}
	}

	for lengthOfStack > 0 {

		postfix = append(postfix, stack[lengthOfStack-1])

		stack = stack[:lengthOfStack-1]

	}

	return postfix , nil 
}
