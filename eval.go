package eval

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// precedence defines the precedence of operators
var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

// Stack is a generic implementation of a stack
type Stack[T any] []T

// append adds a new token to the stack
func (s *Stack[T]) append(token T) {
	*s = append(*s, token)
}

// top returns the top token from the stack
func (s *Stack[T]) top() T {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *Stack[T]) ab() (a T, b T) {
	b = s.top()
	a = s.top()
	return
}

// Output is a generic implementation of a stack
type Output[T any] []T

// append adds a new token to the output
func (o *Output[T]) append(token T) {
	*o = append(*o, token)
}

// top returns the top token from the output
func (o *Output[T]) top() T {
	t := (*o)[len(*o)-1]
	*o = (*o)[:len(*o)-1]
	return t

}

// eval evaluates the given expression in Reverse Polish Notation
func eval(expression string) (int, error) {
	tokens := strings.Split(expression, " ")
	stack := make(Stack[int], 0, 0)

	for _, token := range tokens {
		if v, err := strconv.Atoi(token); err == nil {
			stack = append(stack, v)
		} else if _, ok := precedence[token]; ok && len(token) == 1 {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression: not enough operands")
			}
			a, b := stack.ab()
			switch token {
			case "+":
				stack.append(a + b)
			case "*":
				stack.append(a * b)
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				stack.append(a / b)
			case "-":
				stack.append(a - b)
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		} else {
			return 0, fmt.Errorf("invalid token: %s", token)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression: incorrect number of results on stack")
	}
	return stack[0], nil
}

// rpn converts a given infix expression to Reverse Polish Notation
func rpn(expression string) (string, error) {
	stack := make(Stack[string], 0, 0)
	output := make(Output[string], 0, 0)

	tokens := split(expression)
	if tokens == nil {
		return "", fmt.Errorf("invalid expression: %s", expression)
	}

	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			output.append(token)
		} else if token == "(" {
			stack.append(token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				top := stack.top()
				output.append(top)
			}
			if len(stack) == 0 {
				return "", fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else if p, ok := precedence[token]; ok {
			for len(stack) > 0 && p <= precedence[stack[len(stack)-1]] {
				top := stack.top()
				output.append(top)

			}
			stack.append(token)
		} else {
			return "", fmt.Errorf("unknown token: %s", token)
		}
	}

	for len(stack) > 0 {
		top := stack.top()
		if top == "(" {
			return "", fmt.Errorf("mismatched parentheses")
		}
		output.append(top)
	}

	return strings.Join(output, " "), nil
}

// split splits the given expression into tokens
func split(expression string) []string {
	var tokens []string
	for _, c := range expression {
		if unicode.IsSpace(c) {
			continue
		}
		if c == '(' || c == ')' || c == '+' || c == '-' || c == '*' || c == '/' {
			tokens = append(tokens, string(c))
		} else if unicode.IsDigit(c) {
			if len(tokens) > 0 && unicode.IsDigit(rune(tokens[len(tokens)-1][0])) {
				tokens[len(tokens)-1] += string(c)
			} else {
				tokens = append(tokens, string(c))
			}
		} else {
			return nil
		}
	}
	return tokens
}

// Eval evaluates the given infix expression
func Eval(expression string) (int, error) {
	rpn, err := rpn(expression)
	if err != nil {
		return 0, err
	}
	n, err := eval(rpn)
	if err != nil {
		return 0, err
	}
	return n, nil
}
