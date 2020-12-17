package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("Enter String\n")
	var input string
	n, err := fmt.Scan(&input)
	if n <= 0 {
		fmt.Printf("Scanned 0 or error occured\n")
		fmt.Printf("Err : %s \n", err)
	} else {
		lowerCase := strings.ToLower(input)
		if isFoundain(lowerCase) {
			fmt.Printf("Found ! \n")
		} else {
			fmt.Printf("Not Found ! \n")
		}
	}
}
func isFoundain(str string) bool {

	if strings.Index(str, "i") == 0 && strings.Index(str, "n") == len(str)-1 && strings.IndexAny(str, "a") != -1 {
		return true
	} else {
		return false
	}
}
