package main

import (
	"fmt"
)

func main() {
	var isName bool = true
	isName = 1 == 2
	fmt.Printf("%v %T\n", isName, isName)
	isName = 1 == 1
	fmt.Printf("%v %T\n", isName, isName)
}
