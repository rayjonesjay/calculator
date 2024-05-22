package errorHandle

import (
	"fmt"
)

//InvalidExpression() prints an error if an expression is invalid
func InvalidExpression(invalidToken string) {
	fmt.Printf("invalid token %q\n",invalidToken)
}
