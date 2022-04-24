package main

import (
	"fmt"
	"strconv"
)

var q = 27

func main() {
	i := 42
	//var j float32 = 27
	//var k float64 = 64
	//var l int8 = 23
	//var m int16 = 25
	//var n int32 = 26
	//var o int64 = 29
	//var p string = "Sasi is 28"
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("Global variable %v %T \n", q, q)
	fmt.Println("Public variables are declared as capital letters ex: var Name string= 'name here'")
	fmt.Println("Private variables are declared as small letters ex: var name string= 'name here'")
	var a = float64(i)
	fmt.Printf("%v %T \n", a, a)
	fmt.Println("String conversion are actually done by 'strconv' package")
	fmt.Print("Int to String conv without strconv package: ", string(i))
	fmt.Printf(" -> Type: %T\n", string(i))
	fmt.Println("Int to String conv with strconv package: ", strconv.Itoa(i))
	fmt.Println("All declared variables must be used")
}
