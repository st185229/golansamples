package main

import (
	"fmt"
	"strconv"
)

func main() {

	var f float64

	fmt.Println("Enter a float value : ")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("You have entered : %f \n", f)

	fmt.Println("Alternative output ", strconv.FormatFloat(f, 'f', 6, 64))

}
