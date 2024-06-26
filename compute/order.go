package compute


//Precedence() returns an int indicating the precedence of the operators
func Precedence(op string) int {
	
	switch op {
	case "+", "-":
		return 1
	case "*", "/", "of":
		return 2
	case "^":
		return 3
	}
	return 0
}

