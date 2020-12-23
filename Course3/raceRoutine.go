/*Write two goroutines which have a race condition when executed concurrently.
 Explain what the race condition is and how it can occur.

Submission: Upload your source code for the program along with your written explanation of race conditions.
*/

/*
 The routine 1 and 2 execute due to a delay in the main. But dont finish - So does not execute

 The main start with 2 routines , but they dont finish and they run with different speed due to

*/
/************************************************************Details ***************************
/*The explanation of the race condition and how it can occur is complete, detailed, and clearly written.
The demo shows the race condition when two  operations are making progress at the same time. Concurrent operations are non-deterministic and are therefore unpredictable and extremely hard to reason about. The routine 1 and routine 2 are kick started from the main thread at the same time one soon after another and followed by a sleep in the main thread so that the routine threads can progress.  The routine 1 (fasterRoutine()) has been made a faster one by introducing small latency (sleep 20 ms) and routine 2 (slowerRoutine()) with bigger latency (40 ms) and executed 1000 times.  They get executed but don’t get completed while main thread wait for a short duration . Both of the routines do not return
The race condition becomes apparent when load is applied (delay). The test below applies 1000 requests using 2 goroutines. The expectation is that the counter is incremented for each loop up to 1000 but the results are way way off:
A global variable (shared variable)  routineCounter is an instance of  RoutineExecutions struct is the variable undergoes race condition.  The RoutineExecutions defined with 3 pointer receivers. The first one SetFast(fast) set the execution counter for the faster one and SetSlow(sloe) for the slower one. The third one  GetTotalExecutions() returns the sum of executions
The race condition caused both routine executes different times with different count
-------------------------------------------
Executed routine1 , 24th time
 Executed routine2 , 13th time
 Program ends with total execution 37
-------------------------------------------------
Executed routine1 , 25th time
 Executed routine2 , 13th time
 Program ends with total execution 38
Executed routine2 , 12th time
 Executed routine1 , 23th time
 Executed routine1 , 24th time
 Program ends with total execution 36

******************************************************/

package main

import (
	"fmt"
	"time"
)

/****************************************A global struct and pointer recv**********************************************************************/

// The global variable subjected t= race condition
var routineCounter RoutineExecutions

// RoutineExecutions Structure which holds faster and slower executioncounts
type RoutineExecutions struct {
	countFasterRoutine, countSlowerRoutine int32
}

// SetFast Add the faster routine cunt
func (v *RoutineExecutions) SetFast(fast int32) {
	v.countFasterRoutine += fast
}

// SetSlow Add the slower routine cunt
func (v *RoutineExecutions) SetSlow(slow int32) {
	v.countSlowerRoutine += slow

}

// GetTotalExecutions , Sum of their execution count
func (v *RoutineExecutions) GetTotalExecutions() int32 {
	return v.countFasterRoutine + v.countSlowerRoutine
}

/**************************** Two routines threads which manipulate these values *******************************************************/

// This routine executes a fast loop with smaller sleep time
func fasterRoutine(startupString, routineName string, counter, delay int) {

	fmt.Println(startupString)

	i := 0

	// initializing  for loop
	for i = 1; i <= counter; i++ {

		routineCounter.SetFast(1)

		// each time it write a statement on counter and routine name
		fmt.Printf("\n Executed %s , %dth time", routineName, i)

		// this makes the program sleep for 'delay' ms
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
	// Yoy may never see this statement as main quit
	fmt.Printf("\n Finished executing %s , %dth time", routineName, i)
}

// This routine executes a larger sleep time
func slowerRoutine(startupString, routineName string, counter, delay int) {

	fmt.Println(startupString)
	i := 0
	// initializing a  for loop
	for i = 1; i <= counter; i++ {

		routineCounter.SetSlow(1)
		// prints string and number
		fmt.Printf("\n Executed %s , %dth time", routineName, i)

		// this makes the program sleep for for 'delay' ms
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
	// Yoy may never see this statement
	fmt.Printf("\n Finished executing %s , %dth time", routineName, i)
}

func main() {

	// Simple go program with concurrency

	// Race condition

	// How many time, but never go upto 1000
	const counter = 1000

	//A small delay
	const smallDelay = 20
	go fasterRoutine("*********Welcome to Routine 1 *********", "routine1", counter, smallDelay)

	// Placing the go command in front of the
	// func call simply creates a goroutine
	const bigDelay = 40
	go slowerRoutine("*********Welcome to Routine 2 *********", "routine2", counter, bigDelay)

	//Allow little delay so that routine execute
	time.Sleep(time.Millisecond * 500)

	// The second goroutine, you may think that the
	// program will now run with lightning speed
	// But, both goroutines go to the background
	// It executes partially and quit
	fmt.Printf("\n Program ends with total execution %d", routineCounter.GetTotalExecutions())

	// This statement will now be executed
	// and nothing else will be executed
	// check the output
}
