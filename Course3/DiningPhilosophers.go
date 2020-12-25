/*************************************************************
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
                                        on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
                                     on a line by itself, where <number> is the number of the philosopher.
************************************************************/
package main

import (
	"fmt"
	"sync"
	"time"
)

//Chopstick is a critical section/Mutex
type chopstick struct{ sync.Mutex }

// The structure hold the template of a philosopher
type philosopher struct {
	id                            int
	name                          string
	leftChopstick, rightChopstick *chopstick
}

// Sync wait group , It can be implemented as channel as well
var eatWaitingGroup sync.WaitGroup

// Maximum philospheres are 5
const maxNumberOfPhilosophers = 5

// Eating 3 times
const numberOfCourses = 3

var courses = [...]string{"Starter", "Main", "Dessert"}
var nameOfPhilosophers = [...]string{"Plato", "Aristotle", "Socrates", "Thomas Aquinas", "Immanuel Kant"}

// Goes from thinking to hungry to eating and done eating then starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (p philosopher) eat() {
	for j := 0; j < numberOfCourses; j++ {
		//Pick up sticks means lock both mutex
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		// Acknowledge the start
		//starting to eat <number>
		fmt.Printf("Starting to eat %d(%s) [%s]  \n", p.id+1, nameOfPhilosophers[p.id], courses[j])
		time.Sleep(time.Second)
		//Release mutex
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		// Acknowledge the finish
		//finishing eating <number>
		fmt.Printf("finishing eating %d(%s) [%s]  \n", p.id+1, nameOfPhilosophers[p.id], courses[j])
		time.Sleep(time.Second)
	}
	eatWaitingGroup.Done()
}

func main() {

	// Create forks
	chopsticks := make([]*chopstick, maxNumberOfPhilosophers)
	for i := 0; i < maxNumberOfPhilosophers; i++ {
		chopsticks[i] = new(chopstick)
	}

	// Create philosopher, assign them 2 forks and send them to the dining table
	philosophers := make([]*philosopher, maxNumberOfPhilosophers)
	for i := 0; i < maxNumberOfPhilosophers; i++ {
		philosophers[i] = &philosopher{
			id: i, name: nameOfPhilosophers[i], leftChopstick: chopsticks[i], rightChopstick: chopsticks[(i+1)%maxNumberOfPhilosophers]}
		eatWaitingGroup.Add(1)
		go philosophers[i].eat()

	}
	eatWaitingGroup.Wait()

}
