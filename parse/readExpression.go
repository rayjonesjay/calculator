package parse 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadInput reads a line of input from the user.
func ReadInput() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter expression: ")
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}
