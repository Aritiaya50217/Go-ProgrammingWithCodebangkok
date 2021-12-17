package main

import "fmt"

func main() {
	// array 1D
	x := [3]int{1, 2, 3}
	// เข้าถึง array
	x[0] = 10
	fmt.Printf("%#v\n", x)

	// array 2D
	y := [2][2]int{
		{1, 2},
		{4, 5},
	}
	fmt.Println("แถวที่ 0 ลำดับที่ 1 :", y[0][1])

}
