//
// This will create a race condition between the two routines
// increment and decrement which operate on the same variable
// integer. Depending on the execution order, they will either
// increment the integer several times in a row, decrement several times
// in a row or do switching operations.

package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(integer *int64, wg *sync.WaitGroup) {
	(*integer)++
	fmt.Println(*integer)
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	var integer int64 = 0
	var wg sync.WaitGroup

	for j := 0; j < 10; j++ {
		wg.Add(0)

		go increment(&integer, &wg)
		wg.Wait()

	}

	fmt.Println(integer)
}
