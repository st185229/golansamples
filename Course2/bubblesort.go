/*
*    Write a Bubble Sort program in Go. The program
*   should prompt the user to type in a sequence of up to 10 integers. The program
*    should print the integers out on one line, in sorted order, from least to
*    greatest.
 */

/*
* As part of this program, you should write a function called BubbleSort() which
* takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify
* the slice so that the elements are in sorted order.
 */
/*
* A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent
*  elements in the slice. You should write a Swap() function which performs this operation. Your Swap()
* function should take two arguments, a slice of integers and an index value i which
* indicates a position in the slice. The Swap() function should return nothing, but it should swap
* the contents of the slice in position i with the contents in position i+1.
 */
/*
* Test the program by running it twice and
* testing it with a different sequence of integers each time. The first test
* sequence of integers should be all positive numbers and the second test should
* have at least one negative number. Give 3 points if the program works correctly
* for one test sequence, and give 3 more points if the program works correctly
* for the second test sequence.

* Examine the code. If the code contains a function
* called BubbleSort() which takes a slice of integers as an argument, then
* give another 2 points.If the code
* contains a function called Swap() function which takes two arguments, a slice of
* integers and an index value i, then give another 2 points.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//BubbleSort sorts a slice and change the slice itself
func BubbleSort(slice []int) {
	//Start the loop in reverse order, so the loop will start with length
	//which is equal to the length of input array and then loop until
	//reaches 1
	for i := len(slice); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if slice[j-1] > slice[j] {
				Swap(slice, j-1)
			}
		}
	}
}

// Swap slice in position i with the contents in position i+1.
func Swap(slice []int, i int) {
	intermediate := slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = intermediate

}

//MakeIntSliceFromString return array of integers
func MakeIntSliceFromString(value string) []int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)
	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return ConvertStringSlice2IntegerSlice(results)
}

//ConvertStringSlice2IntegerSlice convert string array to integer array
func ConvertStringSlice2IntegerSlice(scanLine []string) []int {
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

func main() {
	//Declare variable to hold input string
	var inputIntegers string = ""
	//Request user to input
	fmt.Print("Please enter maximum 10 integers delimited with space e.g. 2 212 3001 14 -501 7800 9932 33 45 106 : ")
	//Read
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputIntegers = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
		return
	}
	//Convert the string to integer slice
	intSlice := MakeIntSliceFromString(inputIntegers)
	//Bubble Sort
	BubbleSort(intSlice)
	// Print the sorted array
	fmt.Print("Sorted Array :- ")
	//Print sorted array
	fmt.Println(intSlice)
}
