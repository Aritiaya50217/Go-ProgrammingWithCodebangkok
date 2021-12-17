package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}

	// for
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}

	// while loop
	j := 10
	for j < len(values) {
		fmt.Println(values[j])
		j++
	}

	// for range
	for _, v := range values {
		fmt.Println(v)
	}

}
