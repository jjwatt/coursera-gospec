package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	// fmt.Println("captured: ", input)
	low_input := strings.ToLower(input)
	not_found := false
	switch {
	case strings.Index(low_input, "i") != 0:
		not_found = true
	case strings.LastIndex(low_input, "n") != len(low_input)-1:
		not_found = true
	case strings.Contains(low_input, "a") == false:
		not_found = true
	default:
		fmt.Println("Found!")
		os.Exit(0)
	}
	if not_found == true {
		fmt.Println("Not Found!")
		os.Exit(1)
	}
}
