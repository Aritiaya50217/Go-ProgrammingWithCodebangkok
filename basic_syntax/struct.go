package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Hello() string {
	return "Hello " + p.Name
}

func main() {
	x := Person{
		Name: "Gopher",
		Age:  24,
	}

	fmt.Println(x.Hello())
}
