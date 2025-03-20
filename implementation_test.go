package main

import (
	"fmt"
	"testing"
)

func TestPrefixToInfix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"+ 5 * - 4 2 ^ 3 2", "(5 + ((4 - 2) * (3 ^ 2)))", nil},
		{"* + 1 2 3", "((1 + 2) * 3)", nil},
		{"^ 2 3", "(2 ^ 3)", nil},
		{"+ 1", "", errors.New("недостатньо операндів для оператора")},
		{"1 2 +", "", errors.New("неправильний вираз")},
		{"", "", errors.New("пустий рядок")},
		{"+ a b", "", errors.New("недопустимий символ: a")},
	}

	for _, test := range tests {
		result, err := PrefixToInfix(test.input)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Для виразу %s очікувалась помилка %v, але отримано %v", test.input, test.err, err)
		}
		if result != test.expected {
			t.Errorf("Для виразу %s очікувалось %s, але отримано %s", test.input, test.expected, result)
		}
	}
}

func ExamplePrefixToInfix() {
	expression := "+ 5 * - 4 2 ^ 3 2"
	result, _ := PrefixToInfix(expression)
	fmt.Println(result)
	// Output: (5 + ((4 - 2) * (3 ^ 2)))
}
