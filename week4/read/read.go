/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file has
a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname
for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the
first and last names found in the file. Each struct created will be added to a slice, and after all
lines have been read from the file, your program will have a slice containing one struct for each line
in the file. After reading all lines from the file, your program should iterate through your slice of
structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fullname struct {
	fname string
	lname string
}

func main() {
	var names []*fullname
	// An artificial input source.
	// const input = "Jesse Wattenbarger\nJennifer Wattenbarger\nColette Wattenbarger\n"

	// Get the input filename from the user.
	fmt.Printf("Filename to read: ")
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	inputFileName := inputScanner.Text()

	// Open the file.
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Build the scanner from the artificial input source for testing.
	// scanner := bufio.NewScanner(strings.NewReader(input))

	// Build the scanner from the file to pass the assignment requirements.
	scanner := bufio.NewScanner(inputFile)

	// Iterate through the scanner, splitting up lines and building structs.
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		f := &fullname{fname: tokens[0], lname: tokens[1]}
		names = append(names, f)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// Iterate through the slice of structs and print out fname and lname.
	for _, name := range names {
		fmt.Printf("%v %v\n", name.fname, name.lname)
	}
	os.Exit(0)
}
