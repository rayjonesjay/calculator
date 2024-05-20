package errorHandle

import (
	"fmt"
)

func InvalidExpression(invalidToken string) {
	fmt.Printf("invalid token %q\n",invalidToken)
}
