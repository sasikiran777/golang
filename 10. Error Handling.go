package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// defer always executes at the end of the program

func main() {
	fmt.Println("Execution started")
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
}
