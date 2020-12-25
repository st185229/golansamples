/*

Write a program to sort an array of integers. The program
should partition the array into 4 parts, each of which is
sorted by a different goroutine. Each partition should be
of approximately equal size. Then the main goroutine
should merge the 4 sorted subarrays into one large sorted
array.

The program should prompt the user to input a series of
integers. Each goroutine which sorts Â¼ of the array should
print the subarray that it will sort. When sorting is
complete, the main goroutine should print the entire
sorted list.

*/

package main

import (
	"fmt"
	"sort"
	// "math"
)

func sortArr(arr []int, c chan []int) {
	// Print array to sort
    // Note: %v is the default format
	fmt.Printf("Array to sort: %v\n", arr)

	// Sort array
	sort.Ints(arr)

	// Send to channel
	c <- arr
}

func main() {
	// Init
	c := make(chan []int)
	sortedArr := make([]int, 0)
	startPos := 0

	// Ask for user input
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// Check length of user input to determine chunksize
	arrLen := len(arr)
	chunkSize := arrLen / 4
	leftOver := arrLen % 4

	// Call goroutines to sort
    for i := 0; i < 4; i++ {
		if leftOver != 0 {
			go sortArr(arr[startPos:startPos+chunkSize+1], c)
			leftOver -= 1
			startPos = startPos + chunkSize + 1
		} else {
			go sortArr(arr[startPos:startPos+chunkSize], c)
			startPos = startPos + chunkSize
		}
	}

    // Collect sorting output from goroutines
    // The ellipsis allows unpacking the slice :)
	for i := 0; i < 4; i++ {
		sortedArr = append(sortedArr, <-c...)
	}

    // Sort the compiled array one more time
    sort.Ints(sortedArr)

	// Print compiled output
    fmt.Println("Entire sorted array:")
	fmt.Println(sortedArr)
}