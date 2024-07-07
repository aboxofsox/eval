package eval

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrorInvalidExpression               = fmt.Errorf("invalid expression")
	ErrorMismatchedParentheses           = fmt.Errorf("mismatched parentheses")
	ErrorUnknownToken                    = fmt.Errorf("unknown token")
	ErrorUnknownOperator                 = fmt.Errorf("unknown operator")
	ErrorDivisionByZero                  = fmt.Errorf("division by zero")
	ErrorNotEnoughOperands               = fmt.Errorf("not enough operands")
	ErrorIncorrectNumberOfResultsOnStack = fmt.Errorf("incorrect number of results on stack")
)

// precedence defines the precedence of operators
var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

// stack is a generic implementation of a stack.
type stack[T any] []T

// append adds a new token to the stack.
func (s *stack[T]) append(token T) {
	*s = append(*s, token)
}

// pop returns the top token from the stack.
func (s *stack[T]) pop() T {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

// ab returns th top two tokens from the stack.
func (s *stack[T]) ab() (a T, b T) {
	b = s.pop()
	a = s.pop()
	return
}

// last returns the last token in the stack.
func (s *stack[T]) last() T {
	return (*s)[len(*s)-1]
}

// output is a generic implementation of a stack.
type output[T any] []T

// append adds a new token to the output.
func (o *output[T]) append(token T) {
	*o = append(*o, token)
}

// pop returns the pop token from the output.
func (o *output[T]) pop() T {
	t := (*o)[len(*o)-1]
	*o = (*o)[:len(*o)-1]
	return t

}

// Eval evaluates the given infix expression.
func Eval(expression string) (int, error) {
	exp, err := rpn(expression)
	if err != nil {
		return 0, err
	}
	n, err := eval(exp)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// eval evaluates the given expression in Reverse Polish Notation.
func eval(expression string) (int, error) {
	tokens := strings.Split(expression, " ")
	stack := make(stack[int], 0, 0)

	for _, token := range tokens {
		if v, err := strconv.Atoi(token); err == nil {
			stack.append(v)
		} else if _, ok := precedence[token]; ok && len(token) == 1 {
			if len(stack) < 2 {
				return 0, ErrorNotEnoughOperands
			}
			a, b := stack.ab()
			switch token {
			case "+":
				stack.append(a + b)
			case "*":
				stack.append(a * b)
			case "/":
				if b == 0 {
					return 0, ErrorDivisionByZero
				}
				stack.append(a / b)
			case "-":
				stack.append(a - b)
			default:
				return 0, ErrorUnknownOperator
			}
		} else {
			return 0, ErrorUnknownToken
		}
	}

	if len(stack) != 1 {
		return 0, ErrorIncorrectNumberOfResultsOnStack
	}
	return stack[0], nil
}

// rpn converts a given infix expression to Reverse Polish Notation.
func rpn(expression string) (string, error) {
	stack := make(stack[string], 0, 0)
	output := make(output[string], 0, 0)

	tokens := split(expression)
	if tokens == nil {
		return "", ErrorInvalidExpression
	}

	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			output.append(token)
		} else if token == "(" {
			stack.append(token)
		} else if token == ")" {
			for len(stack) > 0 && stack.last() != "(" {
				top := stack.pop()
				output.append(top)
			}
			if len(stack) == 0 {
				return "", ErrorMismatchedParentheses
			}
			_ = stack.pop()
		} else if p, ok := precedence[token]; ok {
			for len(stack) > 0 && p <= precedence[stack.last()] {
				top := stack.pop()
				output.append(top)

			}
			stack.append(token)
		} else {
			return "", ErrorUnknownToken
		}
	}

	for len(stack) > 0 {
		top := stack.pop()
		if top == "(" {
			return "", ErrorMismatchedParentheses
		}
		output.append(top)
	}

	return strings.Join(output, " "), nil
}

// split splits the given expression into tokens.
// whitespace is ignored.
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
