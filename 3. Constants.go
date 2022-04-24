package main

import "fmt"

const (
	num1 = iota
	num2
	num3
)

const (
	num4 = 8 << 3 // a bit wise left shift(3)
	num5 = 8 | 3  // a bit wise or shift
	num6 = 8 & 3  // a bit wise and operator
	num7 = 8 >> 3 // a bit wise right shift(3)
	num8 = 8 ^ 3  // a bit wise exclusive or
)

func main() {
	var i int64 = 2
	fmt.Printf("%v %T \n", num1, num1)
	fmt.Printf("%v %T \n", num2, num2)
	fmt.Printf("%v %T \n", num3, num3)
	fmt.Printf("%v %T \n", num3+i, num3+i)
	fmt.Printf("%v %T \n", num4, num4)
	fmt.Printf("%v %T \n", num5, num5)
	fmt.Printf("%v %T \n", num6, num6)
	fmt.Printf("%v %T \n", num7, num7)
	fmt.Printf("%v %T \n", num8, num8)
}
