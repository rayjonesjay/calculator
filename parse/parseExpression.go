package parse

import (
	"fmt"
	"strings"
	"calculator/errorHandle"
)

func Parse(expression string) {
	tokens := strings.Fields(expression)

	for _ , token := range tokens {
		if !isNumber(token) && !isOperand(token) {
			errorHandle.InvalidExpression(token)
			DisplayPrompt("")
		}
	}
	fmt.Print(expression)
}
