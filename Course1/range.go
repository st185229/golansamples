package main

import "fmt"

func main() {
	/*	x := []int{4, 8, 5}
		y := -1
		for _, elt := range x {
			if elt > y {
				y = elt
			}
		}
		fmt.Print(y)

		x1 := [...]int {4, 8, 5}
		y1 := x1[0:2]
		z1 := x1[1:3]
		y1[0] = 1
		z1[1] = 3
		fmt.Print(x1)*/
	/*
		x := [...]int{1, 2, 3, 4, 5}
		y := x[0:2]
		z := x[1:4]
		fmt.Print(len(y), cap(y), len(z), cap(z))
	*/

	/*x := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}*/

	/*b := P{"x", -1}
	a := [...]P{P{"a", 10},
		P{"b", 2},
		P{"c", 3}}
	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)*/

	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}

type P struct {
	x string
	y int
}
