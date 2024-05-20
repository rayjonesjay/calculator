package parse 

import (
	"strings"
	"strconv"
)

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func toFloat(token string) float64 {
	float_token, _ := strconv.ParseFloat(token, 64)
	return float_token
}

func isOperand(token string) bool {
	//remove leading and trailing spaces
	token = strings.TrimSpace(token)
	var allAreToken bool = true
	tokens := []string{"+","-","/","*","of","(",")"}
	for _, validToken := range tokens {
		if isNumber(token){
			continue
		}else {
			if validToken != token {
				allAreToken = false
				break
			}
		}
		
	}
	return allAreToken
}