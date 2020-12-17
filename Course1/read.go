/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file
and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and
print the first and last names found in each struct. Submit your source code for the program, “read.go”.
This assignment is worth a total of 10 points:

3 points will be given if a program is written.

2 points will be given if the program compiles correctly.

5 points will be given if test execution is successful and your program: 1. Opens a named text file. 2. Prints all first name/ last name pairs.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type person struct {
	fname string
	lname string
}

func newPerson(first string, last string) *person {
	p := person{fname: first, lname: last}
	return &p
}

// WordCount returns number of words and array of each words
func WordCount(value string) (int, []string) {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results), results
}

func main() {

	fmt.Printf("Enter the file name to parse, press return to confirm : ")
	var input string

	n, err := fmt.Scan(&input)
	if err != nil || n <= 0 {
		fmt.Printf("Err : %s \n", err)
		log.Fatal(err)
	}

	//Open file
	f, err := os.Open(input)
	//Check any error
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//create a slice
	slice := make([]*person, 0, 3)
	//declare a reader
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		count, scanLine := WordCount(line)
		// Every line should exactly 2 words , skip those lines
		counter++
		if count != 2 {
			fmt.Println("-----------------------------")
			fmt.Printf("\n%s - Skipped  line# %d, every line should have 2 words, first name , last name ! \n", scanLine, counter)
			continue
		}
		if len(scanLine[0]) > 20 || len(scanLine[1]) > 20 {
			fmt.Println("-----------------------------")
			fmt.Printf("\n%s  - Skipped  line# %d, the first and last name  max 20 chars long: \n", scanLine, counter)
			continue
		}

		slice = append(slice, newPerson(scanLine[0], scanLine[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("{First Name  ----  Last Name}")
	fmt.Println("-----------------------------")

	for _, name := range slice {
		fmt.Printf("%s\n", *name)
	}

}
