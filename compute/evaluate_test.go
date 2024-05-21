package compute

import (
	"testing"
	"math"
	"calculator/parse"
)

func TestEvaluatePostfix(t *testing.T) {
	//positive infinity 
	Inf := math.Inf(1)
	inputs := map[string]float64{
		"1+2":3,
		"3^3":27,
		"3":3,
		"1+2*(3*3+2)^4":29283,
		"0/12":0,
		"12/0":+Inf,
	}

	for key , value := range inputs{
		tokens, err := parse.Tokenize(key)
	if err != nil {
		return 
	}

	postfixTokens, err := InfixToPostFix(tokens)
	if err != nil {
		return 
	}
	result, _ := EvaluatePostfix(postfixTokens)

	if result != value {
		t.Errorf("got %.2f want %.2f", result, value)
	}
	}
	

}
