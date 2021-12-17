package main

import "fmt"

var y int = 11

func main() {
	// const เปลี่ยนแปลงค่าไม่ได้
	// var เปลี่ยนแปลงค่าได้ มักใช้นอก func
	x := 10
	fmt.Println(x, y)
}
