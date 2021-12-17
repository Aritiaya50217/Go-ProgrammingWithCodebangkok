package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// slice
	x := []int{1, 2, 3, 5, 6, 7}

	// append
	x = append(x, 4)

	// นับ slice
	s := len(x)
	fmt.Printf("%#v\n", x)
	fmt.Printf("Lenght : %#v\n", s)

	// นับตัวอักษร
	name := "abc"
	fmt.Println("len abc : ", utf8.RuneCountInString(name))

	// เลือกช่วงที่ต้องการ
	y := x[1:5]
	fmt.Println(y)
}
