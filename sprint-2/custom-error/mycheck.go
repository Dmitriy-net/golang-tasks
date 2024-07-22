//go:build !solution

package mycheck

import (
	"fmt"
	"strings"
	"unicode"
)

// Определение нового типа для []error
type MyErrors []error

// Реализация метода Error для MyErrors
func (e MyErrors) Error() string {

	var errorStrings []string

	for _, err := range e {
		errorStrings = append(errorStrings, err.Error())
	}

	return strings.Join(errorStrings, ";")
}

// Реализация функции MyCheck
func MyCheck(input string) error {

	var errors MyErrors

	// Проверка на длину строки
	if len(input) >= 20 {
		errors = append(errors, fmt.Errorf("line is too long"))
	}

	// Проверка на наличие цифр
	for _, r := range input {
		if unicode.IsDigit(r) {
			errors = append(errors, fmt.Errorf("found numbers"))
			break
		}
	}

	// Проверка на наличие двух пробелов
	if strings.Count(input, " ") != 2 {
		errors = append(errors, fmt.Errorf("no two spaces"))
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}
