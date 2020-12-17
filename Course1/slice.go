/*
* Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
* The program should be written as a loop.
* Before entering the loop, the program should create an empty integer slice of size (length) 3.
* During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
* The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
* The slice must grow in size to accommodate any number of integers which the user decides to enter.
* The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
 */
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//Read as string
	var input string
	// Start as 3 elements length
	slice := make([]int, 0, 3)

	fmt.Printf("Enter an integers to be added to the slice, enter X in upper case to quit\n")
	for {
		//Read as a string as it need to accommodate X
		n, err := fmt.Scan(&input)
		//Check for any error, if err , give an another try
		if err != nil || n <= 0 {
			fmt.Printf("Invalid input, try another integer, or X in upper case to quit \n")
			fmt.Printf("Err : %s \n", err)
			continue
		} else {
			//Check whether input == "X" and break the loop
			if strings.Compare(input, "X") == 0 {
				break
			}
			//Convert to integer
			s, err := strconv.Atoi(input)
			if err == nil {
				//Valid element to add
				slice = append(slice, s)
			} else {
				// Non integers, another try
				fmt.Printf("Invalid input, try another integer, or X in upper case to quit \n")
			}
		}
		sort.Ints(slice)
		fmt.Printf("Sorted slice is (X in upper case to quit):- ")
		fmt.Println(slice)
	}

}
