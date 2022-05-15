package main

import (
	"fmt"
)

func main() {
	a := 27
	b := a
	fmt.Println(a, b)
	a = 28
	fmt.Println(a, b)
	c := 27
	d := &c
	fmt.Println(&c, d)
	c = 28
	// '*' is de-referencing meaning find the address and pull the value
	fmt.Println(c, *d)

	newArray := [3]int{1, 2, 3}
	fmt.Printf("%v %p %p \n", newArray, &newArray[0], &newArray[1])

	// Golang does not allow to perform pointer arithmetic meaning
	//e := &newArray[1] - 8
	//fmt.Printf("%v \n", e)
	var mystruct *myStruct
	mystruct = &myStruct{
		foo: 29,
	}
	fmt.Println(mystruct)
	var mystruct1 *myStruct
	fmt.Println(mystruct1)
	mystruct1 = new(myStruct)
	mystruct1.foo = 20
	fmt.Println(mystruct1.foo)
}

type myStruct struct {
	foo int
}
