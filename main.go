package main

import (
	"fmt"
	"os"

	"calculator/compute"
	"calculator/parse"
)

var EXIT_SUCCESS = func() { os.Exit(0) }

func main() {
	for {
		expression := parse.ReadInput()

		if expression == "exit" {
			fmt.Println("Exiting...")
			EXIT_SUCCESS()
		}

		tokens, parseTokenError := parse.Tokenize(expression)

		if parseTokenError != nil {
			fmt.Println("Error passing expression.")
			continue
		}

		postfixNotation, postfixNotationConversionError := compute.InfixToPostFix(tokens)
		if postfixNotationConversionError != nil {
			fmt.Println("Error converting to postfix notation")
			continue
		}

		result, evaluatePostfixError := compute.EvaluatePostfix(postfixNotation)

		if evaluatePostfixError != nil {
			fmt.Println("Error evaluating postfix notation")
			continue
		}
	
		fmt.Printf("result of %s = %.2f\n", expression, result)
	}
}
