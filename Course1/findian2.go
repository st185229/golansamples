package main

import (
	buf "bufio"
	format "fmt"
	opr "os"
	str "strings"
)

func main() {

	format.Print("Enter a string: ")
	scanner := buf.NewScanner(opr.Stdin)
	if scanner.Scan() && str.HasPrefix(str.ToLower(scanner.Text()), "i") &&
		str.HasSuffix(str.ToLower(scanner.Text()), "n") &&
		str.Contains(str.ToLower(scanner.Text()), "a") {
		format.Println("Found!")
	} else {
		format.Println("Not Found!")
	}

}
