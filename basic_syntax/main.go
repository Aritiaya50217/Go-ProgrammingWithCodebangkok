package main

import (
	"fmt"
	"gobasic/customer"
)

func main() {
	fmt.Printf("Hello %#v\n", "World")

	// run floder customer
	fmt.Println(customer.Sum(10, 10))
}
