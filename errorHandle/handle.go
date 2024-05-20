package errorHandle

import (
	"fmt"
)

func InvalidExpression(invalidToken string) {
	fmt.Printf("Cannot compute expession: invalid token %q\n",invalidToken)
}
