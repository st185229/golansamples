package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var a string
	fmt.Println("enter your file name, example: name.txt ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a = scanner.Text()

	file, err := os.Open(a)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var sli []Name

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {

		s := strings.Split(scanner.Text(), " ")

		name := Name{
			fname: s[0],
			lname: s[1],
		}
		sli = append(sli, name)
	}

	for i := 0; i < len(sli); i++ {
		fmt.Printf("First Name: %s\tLast Name: %s\n", sli[i].fname, sli[i].lname)
	}

}
