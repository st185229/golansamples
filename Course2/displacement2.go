//closure2.go
package main

import "fmt"

func main() {
	var a, v, s float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

// GenDisplaceFn _
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return a*t*t/2 + v*t + s
	}
}
