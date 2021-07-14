package main

import (
	"fmt"

	encode "go-estudos/first-last/encode"
)

func main() {
	test := []string{"leandro", "tixa", "lombi"}
	fmt.Println(encode.EncodeFirstLast(test))

	// ["leandro", "tixa", "lombi"]
	// ["la", "ti", "lo"]

	// ["cat", "dog", "fish"]
	// ["cg", "dh", "ft"]

	// ["cat", "dog", "fish", "lion"]
	// ["cg", "dh", "fn", "lt"]
}
