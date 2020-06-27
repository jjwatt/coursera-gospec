package main

import "fmt"

func main() {
	var inputFloat float64

	fmt.Printf("Enter a floating point number : ")

	fmt.Scan(&inputFloat)

	var ret int = int(inputFloat)

	fmt.Printf("%d\n", ret)
}
