package main

import "fmt"

func main() {
	point := 50
	if point >= 50 && point <= 100 {
		fmt.Println("pass")
	} else if point >= 20 {
		fmt.Println("not pass")
	}
}
