package parse

import (
	"strconv"
	"unicode"
    "strings"

)

//IsNumber() checks if a token is a number
func IsNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

//Tokenize() convert the expression into a slice of tokens
func Tokenize(input string) ([]string, error) {
  
    input = cleanInput(input)
    var tokens []string
    var numBuffer string

    for _, character := range input {
        if unicode.IsDigit(character) || character == '.' {
            numBuffer += string(character)
        } else {
            if numBuffer != "" {
                tokens = append(tokens, numBuffer)
                numBuffer = ""
            }
            if !unicode.IsSpace(character) {
                tokens = append(tokens, string(character))
            }
        }
    }
    if numBuffer != "" {
        tokens = append(tokens, numBuffer)
    }

    return tokens, nil
}


//cleanInput() makes the string into a valid expression
func cleanInput(input string) string {
    input = strings.ReplaceAll(input, "of","*")
    var result string 
    for i := 0 ; i < len(input); i++{
        result += string(input[i])

        if IsNumber(string(input[i])) &&  i+1 < len(input) && input[i+1] == '(' {
            result += "*"
            input = input+" " 
        }
      
    }
 
    
    return result 
}