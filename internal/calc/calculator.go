package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	token "github.com/JaquesBoeno/MathTokenizer/internal/tokenizer"
)

func Calculate(arr []token.Token) (float64, []token.Token, error) {
	var result float64
	tokens := arr

	for i := 0; i <= len(tokens); i++ {
		if i < len(tokens) {
			if tokens[i].Type == "Operator" && tokens[i].Value == "^" {
				val, err := Power(tokens[i-1], tokens[i+1])
				if err == nil {
					result = val
					tokens = append(append(tokens[0:i-1], token.Token{Type: "Number", Value: fmt.Sprintf("%f", result)}), tokens[i+2:]...)
					i = 0
				} else {
					return 0, tokens, err
				}
			}
		}
	}

	for i := 0; i <= len(tokens); i++ {
		if i < len(tokens) {
			if tokens[i].Type == "Operator" && tokens[i].Value == "*" {
				val, err := Multiplication(tokens[i-1], tokens[i+1])
				if err == nil {
					result = val
					tokens = append(append(tokens[0:i-1], token.Token{Type: "Number", Value: fmt.Sprintf("%f", result)}), tokens[i+2:]...)
					i = 0
				} else {
					return 0, tokens, err
				}
			} else if tokens[i].Type == "Operator" && tokens[i].Value == "/" {
				val, err := Division(tokens[i-1], tokens[i+1])
				if err == nil {
					result = val
					tokens = append(append(tokens[0:i-1], token.Token{Type: "Number", Value: fmt.Sprintf("%f", result)}), tokens[i+2:]...)
					i = 0
				} else {
					return 0, tokens, err
				}
			}
		}
	}

	for i := 0; i <= len(tokens); i++ {
		if i < len(tokens) {
			if tokens[i].Type == "Operator" && tokens[i].Value == "+" {
				val, err := Sum(tokens[i-1], tokens[i+1])
				if err == nil {
					result = val
					tokens = append(append(tokens[0:i-1], token.Token{Type: "Number", Value: fmt.Sprintf("%f", result)}), tokens[i+2:]...)
					i = 0
				} else {
					return 0, tokens, err
				}
			}
		}
	}
	return result, tokens, nil
}

func Sum(left token.Token, right token.Token) (float64, error) {
	rightV, errR := strconv.ParseFloat(right.Value, 64)
	leftV, errL := strconv.ParseFloat(left.Value, 64)

	if left.Type == "Number" && right.Type == "Number" {
		if errL == nil && errR == nil {
			return leftV + rightV, nil
		}
		return 0, errors.New(errL.Error() + errR.Error())
	}

	return 0, errors.New("Os valores tem que numéricos")
}

func Multiplication(left token.Token, right token.Token) (float64, error) {
	rightV, errR := strconv.ParseFloat(right.Value, 64)
	leftV, errL := strconv.ParseFloat(left.Value, 64)

	if left.Type == "Number" && right.Type == "Number" {
		if errL == nil && errR == nil {
			return leftV * rightV, nil
		}
		return 0, errors.New(errL.Error() + errR.Error())
	}

	return 0, errors.New("Os valores tem que numéricos")
}

func Division(left token.Token, right token.Token) (float64, error) {
	rightV, errR := strconv.ParseFloat(right.Value, 64)
	leftV, errL := strconv.ParseFloat(left.Value, 64)

	if left.Type == "Number" && right.Type == "Number" {
		if errL == nil && errR == nil {
			return leftV / rightV, nil
		}
		return 0, errors.New(errL.Error() + errR.Error())
	}

	return 0, errors.New("Os valores tem que numéricos")
}

func Power(left token.Token, right token.Token) (float64, error) {
	rightV, errR := strconv.ParseFloat(right.Value, 64)
	leftV, errL := strconv.ParseFloat(left.Value, 64)

	if left.Type == "Number" && right.Type == "Number" {
		if errL == nil && errR == nil {
			return math.Pow(leftV, rightV), nil
		}
		return 0, errors.New(errL.Error() + errR.Error())
	}

	return 0, errors.New("Os valores tem que numéricos")
}
