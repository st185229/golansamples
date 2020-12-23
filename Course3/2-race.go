//
// This will create a race condition between the two routines
// increment and decrement which operate on the same variable
// integer. Depending on the execution order, they will either
// increment the integer several times in a row, decrement several times
// in a row or do switching operations.

package main

import "fmt"
import "time"

var integer int

func increment() {
  integer++
  fmt.Println(integer)
}

func decrement() {
  integer--
  fmt.Println(integer)
}

func main() {
  integer = 0
  for j:=0; j<10; j++{
    go increment()
    go decrement()
  }
  time.Sleep(10 * time.Second)
}
