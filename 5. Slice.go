package main

import "fmt"

//Slices are similar to arrays but these are natural pointers

func main() {
	var a []int
	a = append(a, 3, 4, 5)
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	newArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(newArray[:])
	fmt.Println(newArray[2:])
	fmt.Println(newArray[:4])
	fmt.Println(newArray[1:2])
	fmt.Println("capacity: ", cap(newArray[1:2]))
	fmt.Println("length: ", len(newArray[1:2]))
	makeExample := make([]int, 3, 100) // make of slice(int, string, ..), length of slice, capacity of slice
	fmt.Println(makeExample)
	makeExample = append(makeExample, []int{1, 2, 3, 4}...)
	fmt.Println(makeExample)
}
