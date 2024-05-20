package parse 

import (
	"bufio"
	"fmt"
	"os"
	"log"
)
func ReadExpression() {
	arguments := os.Args
	if !(len(arguments) == 1 && arguments[0]=="main.go") || (len(arguments) ==1 && arguments[0]=="."){
		log.Fatalf("usage: go run main.go")
		os.Exit(1)
	}
}

func DisplayPrompt(expression string) {
	if expression == "" {
		{}
	}
	fmt.Println("[input expression] : ")
	reader := bufio.NewReader(os.Stdin)
	byteExpression, readError := reader.ReadBytes('\n')
	if readError != nil {
		fmt.Println("Error reading from standard input.")
		os.Exit(1)
	}
	
	Parse(string(byteExpression))
}