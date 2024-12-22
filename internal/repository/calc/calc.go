package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var operators = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

type Calculator struct{}

func moreImportant(op1, op2 string) bool {
	return operators[op1] > operators[op2]
}

func isOperator(s string) bool {
	_, ok := operators[s]
	return ok
}

func SortStationAlgorithm(expression string) ([]string, error) {
	var exp []string
	var stack []string

	var tokens []string
	expression = strings.Replace(expression, " ", "", -1)
	var lit string

	for _, ch := range expression {
		switch ch {
		case '+':
			if lit != "" {
				tokens = append(tokens, lit)
				lit = ""
			}
			tokens = append(tokens, "+")
		case '-':
			if lit != "" {
				tokens = append(tokens, lit)
				lit = ""
			}
			tokens = append(tokens, "-")
		case '*':
			if lit != "" {
				tokens = append(tokens, lit)
				lit = ""
			}
			tokens = append(tokens, "*")
		case '/':
			if lit != "" {
				tokens = append(tokens, lit)
				lit = ""
			}
			tokens = append(tokens, "/")
		case '(':
			if lit != "" {
				tokens = append(tokens, lit)
				lit = ""
			}
			tokens = append(tokens, "(")
		case ')':
			if lit != "" {
				tokens = append(tokens, lit)
				lit = ""
			}
			tokens = append(tokens, ")")
		default:
			lit += string(ch)
		}
	}
	if lit != "" {
		tokens = append(tokens, lit)
		lit = ""
	}

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			exp = append(exp, fmt.Sprintf("%f", num))
		} else if isOperator(token) {

			for len(stack) > 0 && isOperator(stack[len(stack)-1]) && !moreImportant(token, stack[len(stack)-1]) {
				exp = append(exp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				exp = append(exp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("")
			}
			stack = stack[:len(stack)-1]
		} else {
			return nil, errors.New("")
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errors.New("")
		}
		exp = append(exp, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return exp, nil
}

func Repl(rpn []string) (float64, error) {
	var stack []float64
	for _, token := range rpn {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("")
				}
				stack = append(stack, a/b)
			}
		} else {
			return 0, errors.New("")
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("")
	}

	return stack[0], nil
}

func (c Calculator) Calc(expression string) (float64, error) {
	repl, err := SortStationAlgorithm(expression)
	if err != nil {
		return 0, err
	}
	return Repl(repl)
}
