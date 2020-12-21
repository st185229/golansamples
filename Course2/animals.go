//Animals.go
/*********************************************************************************************
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:
1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
----------------------------------------------------
cow	    grass	    walk	             moo
bird	worms	    fly	                 peep
snake	mice	    slither	             hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food,
The Move() method should print the animal’s locomotion, and
The Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.

Submit your Go program source code.
***********************************************************************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//Animal , an abstraction which can act as cow, bird or snake
type Animal struct {
	food       string
	locomotion string
	noise      string
	//Added the type, which makes easier to create a type
	typ string
}

// SetTyp - Type of animal
func (a *Animal) SetTyp(sp string) {
	a.typ = sp
}

// Eat give what does the animal eat
func (a *Animal) Eat() string {
	switch typ := a.typ; typ {
	case "cow":
		a.food = "grass"
	case "bird":
		a.food = "worms"
	case "snake":
		a.food = "mice"
	default:
		a.food = "unknown"
	}
	return a.food
}

// Move - how does this animal move
func (a *Animal) Move() string {
	switch typ := a.typ; typ {
	case "cow":
		a.locomotion = "walk"
	case "bird":
		a.locomotion = "fly"
	case "snake":
		a.locomotion = "slither"
	default:
		a.locomotion = "unknown"
	}
	return a.locomotion
}

// Speak - How do the animal communicate or speak
func (a *Animal) Speak() string {
	switch typ := a.typ; typ {
	case "cow":
		a.noise = "moo"
	case "bird":
		a.noise = "peep"
	case "snake":
		a.noise = "hsss"
	default:
		a.noise = "unknown"
	}
	return a.noise
}

// Tokenize returns collection of words
func tokenize(value string) []string {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return results
}

func stringInSlice(a string, list []string) string {
	for _, b := range list {
		if b == a {
			return a
		}
	}
	return ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(q/Q to exit) > ")
		a := Animal{}
		if scanner.Scan() {
			searchString := scanner.Text()
			words := tokenize(searchString)
			// It should be 2 words
			if len(words) != 2 {
				if len(words) == 1 && strings.ToLower(words[0]) == "q" {
					break
				}
				fmt.Println("Invalid entry, please retry")
				continue
			}
			animaltyp := strings.ToLower(words[0])
			animalAction := strings.ToLower(words[1])
			if animaltyp == stringInSlice(animaltyp, []string{"cow", "bird", "snake"}) && animalAction == stringInSlice(animalAction, []string{"eat", "move", "speak"}) {
				a.SetTyp(animaltyp)
				switch x := animalAction; x {
				case "eat":
					println(a.Eat())
				case "move":
					println(a.Move())
				case "speak":
					println(a.Speak())
				default:
					println("Invalid entry, please retry")
					continue
				}
			} else {
				println("Invalid entry, please retry")
				continue
			}
		} else {
			println("Invalid entry, please retry")
			continue
		}

	}
}
