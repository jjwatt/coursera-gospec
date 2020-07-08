package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var nameAddr map[string]string
	nameAddr = make(map[string]string)
	fmt.Println("Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nameAddr["name"] = scanner.Text()
	fmt.Println("Address: ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nameAddr["address"] = scanner.Text()
	asJ, err := json.Marshal(nameAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(asJ))
	os.Exit(0)
}
