package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 0,3)
	var tmp string
	for {
		fmt.Println("enter an int or X to exit:")
		fmt.Scan(&tmp)
		if tmp == "X"{
			break		
		}
		val,err := strconv.Atoi(tmp)
		if err != nil {
	        	fmt.Println("Wrong Value")
    		} else {
			s=append(s,val)
			sort.Ints(s)
			fmt.Println(s)
		}
	}
}
