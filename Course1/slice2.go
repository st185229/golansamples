package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sl := make([]int, 3)
	cn := 3

	for {
		fmt.Printf("X to Quit: ")
		var n string
		fmt.Scan(&n)

		if n == "X" {
			break
		} else {
			t, _ := strconv.Atoi(n)
			sl = append(sl, t)
			cn++
		}

		sort.Ints(sl[3:cn])

		fmt.Printf("Slicing & Sorting: ")
		for i := 3; i < len(sl); i++ {
			fmt.Printf("%d ", sl[i])
		}
		fmt.Println()
	}
}
