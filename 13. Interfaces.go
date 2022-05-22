package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces intro")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))
	// Integer interface
	myInt := IntCounter(0)
	var incrementer Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(incrementer.Increment())
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Incrementer interface {
	Increment() int
}

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(data)
	return n, err
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
