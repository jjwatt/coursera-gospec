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
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userin := scanner.Text()
		lowin := strings.ToLower(userin)
		if lowin == "X" {
			break
		}
		num, err := strconv.Atoi(lowin)
		if err != nil {
			fmt.Printf("Error: %v", err)
			fmt.Println("Please enter a number or 'X' to quit")
			continue
		}
		li = append(li, num)
		sort.Ints(li)
		fmt.Printf("%v", li)
	}
	fmt.Printf("%v", li)
	os.Exit(0)
}
