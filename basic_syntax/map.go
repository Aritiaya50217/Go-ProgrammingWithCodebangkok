package main

import "fmt"

func main() {
	// map คล้าย dictionary
	countries := map[string]string{}
	// [key] = value
	countries["th"] = "Thailand"
	countries["en"] = "United state"

	fmt.Println(countries["th"])

	// check value
	country, ok := countries["jp"]
	if !ok {
		fmt.Println("no value")
		return
	} else {
		fmt.Println(country)
	}

}
