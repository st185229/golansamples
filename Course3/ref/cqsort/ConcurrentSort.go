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
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

// 4 Go routines
const NumberOfGoRoutines = 4

func main() {

	//Declare variable to hold input string
	var inputIntegers string = ""
	//Request user to input
	fmt.Println("Please a series of integers (12) delimited with space e.g. 2 212 3001 14 -501 7800 9932 33 45 106 100 1: ")
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
	intSlice := MakeIntSliceFromString(inputIntegers)

	// Now create 4 workers and assign the array among them
	assignToWorkers(intSlice)
	// Now it came back from workers
	fmt.Println(intSlice)
}

func assignToWorkers(s []int) {
	if len(s) <= 1 {
		return
	}
	workers := make(chan int, NumberOfGoRoutines-1)
	for i := 0; i < (NumberOfGoRoutines - 1); i++ {
		workers <- 1
	}
	cqsort(s, nil, workers)
}

func cqsort(s []int, done chan int, workers chan int) {
	// report to caller that we're finished
	if done != nil {
		defer func() { done <- 1 }()
	}

	if len(s) <= 1 {
		return
	}
	// since we may use the doneChannel synchronously
	// we need to buffer it so the synchronous code will
	// continue executing and not block waiting for a read
	doneChannel := make(chan int, 1)

	pivotIdx := partition(s)

	fmt.Println("pivotIdx")
	fmt.Println(pivotIdx)

	select {
	case <-workers:
		//Each goroutine take one of the index
		fmt.Println("Worker Routine")
		fmt.Println(s[:pivotIdx+1])

		go cqsort(s[:pivotIdx+1], doneChannel, workers)
	default:
		// if no spare workers, sort synchronously
		cqsort(s[:pivotIdx+1], nil, workers)
		// calling this here as opposed to using the defer
		doneChannel <- 1
	}
	// use the existing goroutine to sort above the pivot
	cqsort(s[pivotIdx+1:], nil, workers)
	// if we used a goroutine we'll need to wait for
	// the async signal on this channel, if not there
	// will already be a value in the channel and it shouldn't block
	<-doneChannel
	return
}

func partition(s []int) (swapIdx int) {
	pivotIdx, pivot := pickPivot(s)
	// swap right-most element and pivot
	s[len(s)-1], s[pivotIdx] = s[pivotIdx], s[len(s)-1]
	// sort elements keeping track of pivot's idx
	for i := 0; i < len(s)-1; i++ {
		if s[i] < pivot {
			s[i], s[swapIdx] = s[swapIdx], s[i]
			swapIdx++
		}
	}
	// swap pivot back to its place and return
	s[swapIdx], s[len(s)-1] = s[len(s)-1], s[swapIdx]
	return
}

func pickPivot(s []int) (pivotIdx int, pivot int) {
	// This should be 1/4th
	pivotIdx = rand.Intn(len(s))
	pivot = s[pivotIdx]
	return
}

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
