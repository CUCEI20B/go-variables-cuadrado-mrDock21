package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var side int = 1

	// Read from console
	_, err := fmt.Scan(&input)

	if err == nil {
		// Scan values from input string
		integer, atoi_err := strconv.Atoi(input)
		if atoi_err == nil {
			side = integer
			fmt.Printf("%d\n", side*side)
		}
	} else {
		fmt.Println(err)
	}
}