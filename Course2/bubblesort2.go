// Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
// The program should print the integers out on one line, in sorted order, from least to greatest.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
// You should write a Swap() function which performs this operation. Your Swap() function should take two arguments,
// a slice of integers and an index value i which indicates a position in the slice.
// The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

package main

import "fmt"

// Swap element at index i with element at index i+1
func Swap(arr []int, i int) {
	temp := arr[i]

	arr[i] = arr[i+1]
	arr[i+1] = temp
}

// BubbleSort : Sort the elements in slice arr, using the bubble sort algorithm
func BubbleSort(arr []int) {
	length := len(arr) - 1

	// Loop as many times as the number of elements in the slice
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			// if current element at j is greater that its adjacent element, swap them
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	arr := make([]int, 10)

	fmt.Println("Please enter 10 integers : ")

	var input int
	for i := 0; i < 10; i++ {
		fmt.Scan(&input)
		arr[i] = input
	}

	fmt.Println("Input : ", arr)

	BubbleSort(arr)
	fmt.Println("Result after sorting : ", arr)
}
