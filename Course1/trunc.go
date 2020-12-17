// Created by ST as the part of coursera
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//Declare variable to hold input value
	var inputValue string = ""
	//Request user to input
	fmt.Println("Please enter a floating point value")
	fmt.Scanf("%s", &inputValue)
	if s, err := strconv.ParseFloat(inputValue, 64); err == nil {
		s = math.Trunc(s)
		fmt.Printf("%.0f \n", s)
	}

}
