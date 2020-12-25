/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

/*
Peer reviewer: Please test the program by entering a series of 12 unsorted positive integers.


5 pts
The program prints a sorted series of the 12 integers.


0 pts
The program does not work (no sorted integers produced).

Peer reviewer: Please review the code and observe the program output.


5 pts
There are four goroutines and each one prints a set of the three array elements that it is sorting.


0 pts
There aren't four separate goroutines and/or there is no printout of the three array elements.

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"sync"
)

//MakeIntSliceFromString return array of integers
func makeNSplitIntSliceFromString(value string, batches int) [][]int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)
	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return convertStringSlice2IntegerSlice(results, batches)
}

//ConvertStringSlice2IntegerSlice convert string array to integer array of 4( batches)
func convertStringSlice2IntegerSlice(scanLine []string, batches int) [][]int {
	intSlice := make([]int, 0, 10)
	for _, i := range scanLine {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("%s invalid skipped, ", i)
			//Skip the invalid entry
			continue
		}
		intSlice = append(intSlice, j)
	}
	var divided [][]int
	chunkSize := (len(intSlice) + batches - 1) / batches
	for i := 0; i < len(intSlice); i += chunkSize {
		end := i + chunkSize

		if end > len(intSlice) {
			end = len(intSlice)
		}

		divided = append(divided, intSlice[i:end])
	}
	return divided
}

func sortSlice(slice []int, wg *sync.WaitGroup, counter int) {
	fmt.Printf("\nRoutine %d starting\n", counter)
	sort.Ints(slice)
	wg.Done()

}
func concatAppend(slices [][]int) []int {
	var tmp []int
	for _, s := range slices {
		tmp = append(tmp, s...)
	}
	return tmp
}

func main() {
	//Declare variable to hold input string
	var inputIntegers string = ""
	//Request user to input
	fmt.Println("Please a series of integers delimited with space e.g. 2 212 3001 14 -501 7800 9932 33 45 106 : ")
	//Read
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputIntegers = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error, please try again")
		return
	}
	const maxNumberOfRotines = 4

	intSlice := makeNSplitIntSliceFromString(inputIntegers, maxNumberOfRotines)

	var wg sync.WaitGroup

	fmt.Println("**********************Unsorted split slices to be delegated to routine*************************")
	counter := 0
	for _, i := range intSlice {
		counter++
		fmt.Printf("\n The slice to be sorted by routine %d is :-  %v", counter, i)
		wg.Add(1)
		//It create maxNumberOfRotines = 4

		go sortSlice(i, &wg, counter)

	}
	wg.Wait()
	fmt.Println("")
	fmt.Println("****************individual slices are sorted by routines********************************************")
	//Individual slices are sorted
	for _, i := range intSlice {
		fmt.Printf("\n The slices sorted by routines  %v", i)
	}
	fmt.Println("")

	//Concatenated slices
	concatenatedSlice := concatAppend(intSlice)

	fmt.Printf("\n The combined slice before sorting   %v", concatenatedSlice)

	wg.Add(1)
	go sortSlice(concatenatedSlice, &wg, counter)

	fmt.Printf("\n The combined slice after sorting   %v", concatenatedSlice)
	fmt.Println("")
	fmt.Println("****************The END************************")

}
