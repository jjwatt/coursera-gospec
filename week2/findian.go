package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scan: %v\n", err)
	}
	low_input := strings.ToLower(input)
	switch {
	case strings.Index(low_input, "i") != 0:
		fmt.Println("Not Found!")
		os.Exit(1)
	case strings.LastIndex(low_input, "n") != len(low_input) - 1:
		fmt.Println("Not Found!")
		os.Exit(1)
	case strings.Contains(low_input, "a") == false:
		fmt.Println("Not Found!")
		os.Exit(1)
	default:
		fmt.Println("Found!")
		os.Exit(0)
	}
}
