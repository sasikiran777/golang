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

	a := "start"
	defer fmt.Println(a) // At this point an is already inserted into defer,
	// so any value changes after this will not affect this a value
	a = "end"

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Hello Go!"))
	//})
	//err = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err.Error())
	//}

	// order of execution
	//fmt.Println("Started")
	//defer fmt.Println("defer")
	//panic("Something went wrong")
	//fmt.Println("Ended")

	// panicker fails but still end statement will execute
	fmt.Println("star")
	panicker()
	fmt.Println("end")
}

func panicker() {
	defer func() {
		if w := recover(); w != nil {
			log.Println("Err", w)
		}
	}()
	panic("Something went wrong")
}
