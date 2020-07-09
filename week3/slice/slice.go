package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	li := make([]int, 0, 3)
	for {
		fmt.Println("Please enter a number or 'X' to quit")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userin := scanner.Text()
		lowin := strings.ToLower(userin)
		if lowin == "x" {
			break
		}
		num, err := strconv.Atoi(lowin)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		li = append(li, num)
		sort.Ints(li)
		fmt.Printf("%v\n", li)
	}
	fmt.Printf("%v\n", li)
	os.Exit(0)
}
