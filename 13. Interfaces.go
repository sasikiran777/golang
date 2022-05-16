package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces intro")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(data)
	return n, err
}
