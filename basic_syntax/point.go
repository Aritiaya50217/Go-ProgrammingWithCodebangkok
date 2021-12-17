package main

import "fmt"

func sum(result *int) {
	a := 10
	b := 10
	*result = a + b
}

func main() {
	x := 10
	y := x

	y = 20
	fmt.Printf("value x :%#v\nvalue y :%#v\n", x, y)

	// pointer
	fmt.Println("Address x :", &x)
	fmt.Println("Address y :", &y)

	// ให้ z ไปอยู่ที่ตำแหน่ง x
	var z *int
	z = &x
	fmt.Println("Address z :", z)

	// value z
	fmt.Println("value z :", *z)

	// func sum()
	var s int
	sum(&s)
	fmt.Println(s)

}
