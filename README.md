
```markdown
## Project Name: Calculator

### Program Name: `main.go`

This is a simple calculator program that respects the rules of **BODMAS** (Brackets, Orders (i.e., powers and square roots, etc.), Division and Multiplication, Addition and Subtraction) or **PEMDAS** (Parentheses, Exponents, Multiplication and Division, Addition and Subtraction).

### Algorithm Used: Shunting Yard Algorithm

The calculator uses the Shunting Yard algorithm to convert expressions from **Infix Notation** (e.g., `1 + 2 * 3`) to **Reverse Polish Notation (RPN)** or **Postfix Notation** (e.g., `1 2 3 * +`). This conversion simplifies the computation process.

#### How the Shunting Yard Algorithm Works:

1. Initialize an empty stack for operators and an empty list for the output (postfix notation).
2. Read the tokens of the input expression one by one:
   - If the token is a number, append it to the output list.
   - If the token is an operator, pop operators from the stack to the output list until the top of the stack has an operator of lower precedence or the stack is empty, then push the current operator onto the stack.
   - If the token is a left parenthesis `(`, push it onto the stack.
   - If the token is a right parenthesis `)`, pop operators from the stack to the output list until a left parenthesis is encountered on the stack. Remove the left parenthesis.
3. After reading all tokens, pop any remaining operators on the stack to the output list.

#### Step-by-Step Example: Converting `1 + 2 * 3`

1. **Input:** `1 + 2 * 3`
2. **Initialize:**
   - Stack: `[]`
   - Postfix List: `[]`
3. **Step 1:**
   - Token: `1` (number)
   - Postfix List: `[1]`
4. **Step 2:**
   - Token: `+` (operator)
   - Stack: `["+"]`
5. **Step 3:**
   - Token: `2` (number)
   - Postfix List: `[1, 2]`
6. **Step 4:**
   - Token: `*` (operator)
   - Stack: `["+", "*"]` (since `*` has higher precedence than `+`)
7. **Step 5:**
   - Token: `3` (number)
   - Postfix List: `[1, 2, 3]`
8. **Final Steps:**
   - Pop all operators from the stack to the Postfix List.
   - Postfix List: `[1, 2, 3, "*", "+"]`

### Usage:

To use the calculator, simply run the program and enter an expression in infix notation when prompted. The program will convert the expression to postfix notation using the Shunting Yard algorithm and then evaluate it.

1. Run the program:
   ```sh
   go run main.go
