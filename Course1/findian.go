// Created by ST as the part of coursera
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	//Declare variable to hold input string
	var searchString string = ""
	//Request user to input
	fmt.Print("Please enter a string to search: ")
	//read string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		searchString = scanner.Text()
	}
	//Quickest way is to use a regex
	var validStringPattern = regexp.MustCompile(`^(i|I).*[aA].*(n|N)$`)

	if validStringPattern.MatchString(searchString) == true {
		fmt.Println("Found")

	} else {
		fmt.Println("Not Found!")
	}

}
