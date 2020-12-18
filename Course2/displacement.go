// Displacement
/*
*
*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s =½ a t2 + vot + so

1) Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFloat Function which read a float value
func ReadFloat(message string) float64 {

	stdin := bufio.NewReader(os.Stdin)
	floatValue := 0.0
	fmt.Println(message)
	for {
		_, err := fmt.Fscan(stdin, &floatValue)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
		fmt.Printf("Invalid input! %s\n", message)
	}
	return floatValue
}

// GenDisplaceFn take acceleration a, initial velocity vo and initial displacement as args & returns computes displacement as a function of time
func GenDisplaceFn(so float64, vot float64, a float64) func(t float64) float64 {

	return func(t float64) float64 {
		//s =½ a t2 + vot + so
		return (1.0/2.0)*a*t*t + vot*t + so
	}
}

// main entry point
func main() {
	so := ReadFloat("Enter value for Initial Displacement (m)")
	vot := ReadFloat("Enter value for Initial Velocity (m/s)")
	a := ReadFloat("Enter value for Acceleration(m/s^2)")
	t := ReadFloat("Enter value for time(s)")
	fmt.Println(so, vot, a, t)
	displacementFunction := GenDisplaceFn(so, vot, a)
	fmt.Println(displacementFunction(t))

}
