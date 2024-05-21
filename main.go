package main

import (
	"fmt"
	"os"

	"calculator/compute"
	"calculator/parse"
)

var EXIT_SUCCESS = func() { os.Exit(0) }

func main() {

	fmt.Fprintf(os.Stdout, "****type exit to quit****\n")
	// fmt.Printf("If resut exceeds [%f] +Inf will be printed.\n",maxFloat)
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
			fmt.Println("Error evaluating expression")
			continue
		}

		if compute.HasFractionPart(result) {
			fmt.Printf("result > %.2f\n", result)
		} else {
			fmt.Printf("result > %.0f\n", result)
		}

	}
}
