package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name  string `required max:"100"`
	phone int
	age   int
}

type Job struct {
	Person
	title string
}

func main() {
	person := Person{
		name:  "Sasi",
		phone: 8500035146,
		age:   28,
	}
	fmt.Println(person)
	newStruct := struct{ name string }{name: "Sasi"}
	copyStruct := newStruct // Struct are not natural pointers to achieve pointers use &
	copyStruct.name = "kiran"
	fmt.Println(newStruct)
	fmt.Println(copyStruct)
	copyStruct1 := &newStruct
	copyStruct1.name = "kiran"
	fmt.Println(newStruct)
	fmt.Println(copyStruct1)
	job := Job{}
	job.name = "Bash"
	job.age = 37
	job.phone = 1234567890
	job.title = "Sr. UI/UX"
	fmt.Println(job.Person)
	fmt.Println(job.Person.name + "\n" + job.title)
	relection := reflect.TypeOf(Person{})
	tag, _ := relection.FieldByName("name")
	fmt.Println(tag.Tag)
}
