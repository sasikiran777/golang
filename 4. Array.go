package main

import "fmt"

func main() {
	//var grades [3]int // dynamic arrays
	//grades[0] = 97
	//grades[1] = 98
	//grades[2] = 96
	grades := [3]int{97, 98, 96} // static arrays
	fmt.Printf("%v %T \n", grades, grades)
	students := [...]string{"sasi", "sam", "bas"}
	fmt.Println(len(students))
	students[2] = "bash"
	fmt.Println(students)
	identityMatrix := [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	c := [...]int{1, 2, 3}
	d := &c // d pointing to c, '&' represents pointers in goang
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)
}
