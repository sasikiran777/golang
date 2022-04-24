package main

import (
	"fmt"
	"strconv"
)

func main() {
	mapExample := map[string]string{
		"name":  "Sasi",
		"age":   "28",
		"phone": strconv.Itoa(8500035146),
	}
	fmt.Println(mapExample)
	testMap := mapExample // maps are natural pointers
	testMap["gender"] = "male"
	fmt.Println(testMap, mapExample)
	delete(testMap, "gender")
	fmt.Println(testMap, mapExample)
	_, ok := testMap["gender"] // _ is written only syntax no need to read
	fmt.Println("is key exist:", ok)
	makeMapExample := make(map[string]string, 10)
	fmt.Println(makeMapExample, len(makeMapExample))
}
