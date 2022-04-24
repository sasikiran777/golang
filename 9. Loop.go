package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 0; i <= 5; i, j = i+1, j+1 {
		fmt.Println("i= ", i)
		fmt.Println("j= ", j)
	}
	a := 0
	for ; a <= 5; a++ {
		fmt.Println("a= ", a)
	}
	b := 0
	for b <= 5 {
		fmt.Println("b= ", b)
		b++
		break
	}
Loop:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if i > 3 {
				fmt.Println(i)
				break Loop
			}
		}
	}

	c := []int{1, 2, 3, 4}
	for k, v := range c {
		fmt.Println(k, v)
	}
}
