package utils

import (
	"regexp"
)

func IsDigit(ch string) bool {
	return regexp.MustCompile(`[\d,.]`).MatchString(ch)
}

func IsOperation(ch string) bool {
	return regexp.MustCompile(`\+|-|\*|\/|\^`).MatchString(ch)
}

func IsLeftParenthesis(ch string) bool {
	return ch == "("
}

func IsRightParenthesis(ch string) bool {
	return ch == ")"
}
