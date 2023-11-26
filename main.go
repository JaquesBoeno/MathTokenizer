package main

import (
	"fmt"

	token "github.com/JaquesBoeno/MathTokenizer/internal/tokenizer"
)

func main() {
	fmt.Println(token.Tokenizer("8+9-2"))
}
