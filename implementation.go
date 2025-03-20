package main

import (
	"errors"
	"strings"
	"unicode"
)

// Функція для перетворення префіксного виразу в інфіксний
func PrefixToInfix(expression string) (string, error) {
	if len(expression) == 0 {
		return "", errors.New("пустий рядок")
	}

	tokens := strings.Fields(expression)
	stack := make([]string, 0)

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("недостатньо операндів для оператора")
			}

			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			infix := "(" + operand1 + " " + token + " " + operand2 + ")"
			stack = append(stack, infix)
		} else if isNumber(token) {
			stack = append(stack, token)
		} else {
			return "", errors.New("недопустимий символ: " + token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("неправильний вираз")
	}

	return stack[0], nil
}

// Функція для перевірки, чи є токен оператором
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

// Функція для перевірки, чи є токен числом
func isNumber(token string) bool {
	for _, char := range token {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
