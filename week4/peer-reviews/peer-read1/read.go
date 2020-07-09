package main
import (
	"fmt"
	"os"
	"strings"
	"bufio"
)
type name struct{
	fname string
	lname string
}
func main() {
	var filename string
	fmt.Println("Enter File Name")
	fmt.Scanln(&filename)
	file, _ := os.Open(filename)

	defer file.Close()
	names := make([]name,0)

	s := bufio.NewScanner(file)
	for s.Scan(){

		s := strings.Split(s.Text()," ")
		a := name{fname:string(s[0]),lname:string(s[1])}
		names = append(names,a)
	}

	for _, v := range names {
		fmt.Printf("First name: %s\n Last name: %s\n",v.fname,v.lname)
	}
}