package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// fmt.Println("Hello, World!"
	fmt.Println("Enter a float: ")
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scan: %v\n", err)
	}
	input_trunc := math.Trunc(input)
	fmt.Println(input_trunc)
}
