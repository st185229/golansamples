package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Printf("enter a floating number : ")
	fmt.Scan(&number)
	fmt.Printf(" Here is int number : %d\n", int64(number))
}
