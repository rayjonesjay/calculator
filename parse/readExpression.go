package parse 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)
func ReadExpression() {
	arguments := os.Args
	if !(len(arguments) == 1 || arguments[0]=="main.go"){
		log.Fatalf("usage: go run main.go")
		os.Exit(1)
	}
	
	DisplayPrompt()
}

func printMessage(){
	message := ("type quit to exit")
	fmt.Println(message)
}

func DisplayPrompt() {
	
	numberOfTimesProgramIsRan := 0
	
	fmt.Println(numberOfTimesProgramIsRan)
	if numberOfTimesProgramIsRan == 1{
		printMessage()
		numberOfTimesProgramIsRan++
	}
	
	write := bufio.NewWriter(os.Stdout)
	_, writeErr := write.WriteString("> ")
	if writeErr != nil {
		_ = fmt.Errorf("error writing to stdout: %v",writeErr)
	}

	flushError := write.Flush()
	if flushError != nil {
		_ = fmt.Errorf("error flushing to stdout: %v",flushError)
	}

	reader := bufio.NewReader(os.Stdin)
	byteExpression, readError := reader.ReadBytes('\n')
	stringExpression := string(byteExpression)
	if stringExpression == "" {
		println("no expression passed")
	}

	if strings.EqualFold(strings.TrimSpace(string(byteExpression)),"quit"){
		os.Exit(0)
	}

	if readError != nil {
		fmt.Println("Error reading from standard input.")
		os.Exit(1)
	}
	
	Parse(string(byteExpression))
}