package main

import "fmt"

func sum(a int, b int) (string, int) {
	return "Result :", a + b
}

// (int,int)int
func add(a, b int) int {
	return a + b
}

// (int , int) int
func sub(a, b int) int {
	return a - b
}

// (string)string
func hello(name string) string {
	return "hello" + name
}

func cal(f func(int, int) int) {
	number := f(50, 10)
	fmt.Println(number)
}

func test(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	x, s := sum(10, 100)
	fmt.Println(x, s)

	cal(add)

	t := test(1, 2, 3, 4)
	fmt.Println(t)

}
