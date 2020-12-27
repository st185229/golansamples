/*
a) The program asks the user to enter a number of integers
b) They are validated

The program prints a sorted series of the 12 integers.
The user inputs are converted into a slice of integers

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const maxGoRoutines = 4

func main() {

	//Declare variable to hold input string
	var inputIntegers string = ""
	//Request user to input
	fmt.Println("Please a series of integers (e.g. 12) delimited with space e.g. 2 212 3001 14 -501 7800 9932 33 45 106 100 1: ")
	//Read
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputIntegers = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error, please try again")
		return
	}
	// The intSlice holds the entire intiger array
	intSlice := makeIntSliceFromString(inputIntegers)
	arrSize := len(intSlice)
	// Make channel
	threadChannel := make(chan []int)
	sortGroup := arrSize / maxGoRoutines
	nosBeyondMultiples4 := arrSize % maxGoRoutines
	sortedArr := make([]int, 0)

	// Call goroutines to sort
	movingPointer := 0
	for i := 0; i < 4; i++ {
		if nosBeyondMultiples4 != 0 {
			go sortSlice(intSlice[movingPointer:movingPointer+sortGroup+1], threadChannel)
			nosBeyondMultiples4--
			movingPointer = movingPointer + sortGroup + 1
		} else {
			go sortSlice(intSlice[movingPointer:movingPointer+sortGroup], threadChannel)
			movingPointer = movingPointer + sortGroup
		}
	}

	for i := 0; i < maxGoRoutines; i++ {
		sortedArr = append(sortedArr, <-threadChannel...)
	}

	// Sort the compiled array one more time
	sortSlice(sortedArr, nil)

	// Print compiled output
	fmt.Println("Sorted array in whole is printed below")
	fmt.Println(sortedArr)
}

// This creates integer slice
func makeIntSliceFromString(value string) []int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)
	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return convertStringSlice2IntegerSlice(results)
}

//ConvertStringSlice2IntegerSlice convert string array to integer array
func convertStringSlice2IntegerSlice(scanLine []string) []int {
	intSlice := make([]int, 0, 10)
	for _, i := range scanLine {
		j, err := strconv.Atoi(i)
		if err != nil {

			fmt.Println("Skipping invalid input error :- ")
			fmt.Println(err)
			//Skip the invalid entry
			continue
		}
		intSlice = append(intSlice, j)
	}
	return intSlice
}
func sortSlice(slice []int, c chan []int) {

	fmt.Printf("Partitioned array before sorting : %v\n", slice)
	sort.Ints(slice)
	if c != nil {
		c <- slice
	}
}
