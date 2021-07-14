package main

import (
	"fmt"

	encode "go-estudos/first-last/encode"
)

func main() {
	test := []string{"Leandro", "Tixa", "Lombi"}
	fmt.Println(encode.EncodeFirstLast(test))

}
