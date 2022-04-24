package main

import (
	"fmt"
	"math"
)

func main() {
	mapValue := map[string]string{
		"name":  "Sasi",
		"phone": "8500035146",
		"age":   "28",
	}
	firstNumber := 36
	secondNumber := 45
	if _, ok := mapValue["t"]; ok {
		fmt.Println("hi")
	}
	if _, ok := mapValue["phone"]; ok {
		fmt.Println(mapValue["phone"])
	}
	if len(mapValue["name"]) > 0 {
		fmt.Println("Got into name: ", mapValue["name"])
	}
	fmt.Println(firstNumber >= secondNumber, secondNumber >= firstNumber, firstNumber == secondNumber)
	if math.Pow(math.Sqrt(0.1), 2) == 0.1 {
		fmt.Println(math.Pow(math.Sqrt(0.1), 2))
	}

	if math.Abs(0.123/math.Pow(math.Sqrt(0.123), 2)-1) < 0.001 {
		fmt.Println(math.Abs(0.123/math.Pow(math.Sqrt(0.123), 2) - 1))
	}
}
