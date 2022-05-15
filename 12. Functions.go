package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println(newFunc("World"))
	name := "Sasi"
	greeting := "Hi"
	greet(&greeting, &name)
	fmt.Println(greeting, name)
	fmt.Println(sum(1, 2, 3, 4, 5))
	s := sum1(6, 7, 8, 9, 10)
	fmt.Println("Pointer style Example", *s)
	fmt.Println(sum2(11, 12, 13, 14, 15))
	div, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(div)
	func() {
		fmt.Println("This is anonymous function")
	}()
	var divideAnonymous func(float64, float64) (float64, error)
	divideAnonymous = func(a float64, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("divide by zero error")
		}
		return a / b, nil
	}
	d, err := divideAnonymous(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	// Methods are similar to functions in go but only difference is they operate in known context(type keyword)
	//
	// Example there is no need to pass greeter as explicit param for greetings method greeter object(g)
	// is passed down to greetings automatically
	g := greeter{
		name:     "Sasi!!",
		greeting: "Hello",
	}
	g.greetings("addtional param")
	g1 := greeter{
		name:     "Sasi!!",
		greeting: "Hello",
	}
	fmt.Println("Name before:", g1.name)
	g1.greetings1("addtional param")
	fmt.Println("Name after:", g1.name)
}

func newFunc(variable string) string {
	fmt.Println(variable)
	return "name"
}

func greet(greetings, name *string) {
	fmt.Println(*greetings, *name)
	*name = "Sasi Kiran"
}

func sum(value ...int) int {
	fmt.Println(value)
	sum := 0
	for _, v := range value {
		sum += v
	}
	return sum
}

func sum1(value ...int) *int {
	fmt.Println(value)
	sum := 0
	for _, v := range value {
		sum += v
	}
	return &sum
}

func sum2(value ...int) (sum int) {
	fmt.Println(value)
	for _, v := range value {
		sum += v
	}
	return
}

func divide(a float64, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("divide by zero error")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greetings(addParam string) {
	fmt.Println(g.greeting, g.name, addParam)
}

func (g *greeter) greetings1(addParam string) {
	fmt.Println(g.greeting, g.name, addParam)
	g.name = "Sasi Kiran"
}
