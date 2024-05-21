package parse

import (
	"strconv"
	"unicode"
)

func IsNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func Tokenize(input string) ([]string, error) {
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
