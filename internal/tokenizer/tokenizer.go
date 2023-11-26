package token

import (
	"regexp"
	"strings"

	"github.com/JaquesBoeno/MathTokenizer/internal/utils"
)

type Token struct {
	Type  string
	Value string
}

func Tokenizer(str string) []Token {
	var results []Token
	str = regexp.MustCompile(`\s+`).ReplaceAllString(str, "")

	tokens := strings.Split(str, "")
	var currentNumber string

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "-" {
			if currentNumber != "" {
				results = append(results, Token{"Number", currentNumber})
				results = append(results, Token{"Operator", "+"})

				currentNumber = ""
			}

			currentNumber += tokens[i]
		} else if utils.IsDigit(tokens[i]) {
			currentNumber += tokens[i]
		} else {
			if currentNumber != "" {
				results = append(results, Token{"Number", currentNumber})
				currentNumber = ""
			}
			if utils.IsOperation(tokens[i]) {
				results = append(results, Token{"Operator", tokens[i]})
			} else if utils.IsLeftParenthesis(tokens[i]) {
				results = append(results, Token{"Left Parenthesis", tokens[i]})
			} else if utils.IsRightParenthesis(tokens[i]) {
				results = append(results, Token{"Right Parenthesis", tokens[i]})
			}
		}
	}

	if currentNumber != "" {
		results = append(results, Token{"Number", currentNumber})
	}

	return results
}
