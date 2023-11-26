package main

import (
	"fmt"

	"github.com/JaquesBoeno/MathTokenizer/internal/calc"
	token "github.com/JaquesBoeno/MathTokenizer/internal/tokenizer"
)

func main() {
	fmt.Println(calc.Calculate(token.Tokenizer("2+9*2^3/2-1")))
}
