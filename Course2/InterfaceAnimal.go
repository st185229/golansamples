//Interface.go
/************************************************************************************************************************************************
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command,
the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.
------------------------------------------------------------------------------------------------------------------------------------------------
Animal		Food eaten		Locomotion method		Spoken sound
cow			grass			walk					moo
bird		worms			fly						peep
snake		mice			slither					hsss
------------------------------------------------------------------------------------------------------------------------------------------------
Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information requested about the animal, either
 “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and
Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
********************************************************************************************************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Animal which define the behavior
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Get the Animal types supported by this program
func getSupportedAnimalType() []string {
	return []string{"cow", "bird", "snake"}
}

// What are the activities
func getSupportedActivity() []string {
	return []string{"eat", "move", "speak"}
}

// What arguments
func getSupportedArguments() []string {
	return []string{"newanimal", "query"}
}

// Cow type of animal
type Cow struct {
}

//Eat Cow eat
func (f Cow) Eat() {
	fmt.Println("grass")
}

// Move of Cow
func (f Cow) Move() {
	fmt.Println("walk")
}

//Speak cow Speaks
func (f Cow) Speak() {
	fmt.Println("moo")
}

// Bird type of animal
type Bird struct {
}

//Eat Bird eat
func (f Bird) Eat() {
	fmt.Println("worms")
}

// Move of Bird
func (f Bird) Move() {
	fmt.Println("fly")
}

//Speak Bird Speaks
func (f Bird) Speak() {
	fmt.Println("peep")
}

// Snake animal
type Snake struct {
}

//Eat Snake eat
func (f Snake) Eat() {
	fmt.Println("mice")
}

// Move of Snake
func (f Snake) Move() {
	fmt.Println("slither")
}

//Speak Snake Speaks
func (f Snake) Speak() {
	fmt.Println("hsss")
}

// UserInputType is custom type to gather user inputs
type UserInputType string

// Tokenize returns collection of words
func (u UserInputType) Tokenize() []string {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)
	// Find all matches and return count.
	results := re.FindAllString(string(u), -1)
	return results
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ToLower lower case user input
func (u UserInputType) ToLower() UserInputType {

	return UserInputType(strings.ToLower(string(u)))

}

func validateInputSlice(words []string) (bool, bool, string) {
	quit := false
	invalid := false
	action := ""

	if len(words) != 3 {
		if len(words) == 1 && words[0] == "q" {
			quit = true
		}
		invalid = true
	} else {
		action = strings.ToLower(words[0])
		if !stringInSlice(action, getSupportedArguments()) {
			invalid = true
		}
	}
	return quit, invalid, action
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	quit := false
	invalid := false
	action := ""
	const errorStatement = "Invalid entry, please retry, please enter three words [newanimal|query]  [animal_name]  [type|activity] "
	const errorNotFound = "Animal not found use newanimal   [animal_name]  [type] to create a new one"
	userAnimalCollection := make(map[string]Animal)
	for {
		fmt.Print("\n(q/Q to exit) > ")
		words := make([]string, 0, 3)
		if scanner.Scan() {
			words = UserInputType(scanner.Text()).ToLower().Tokenize()
			quit, invalid, action = validateInputSlice(words)
			if quit {
				return
			}
			if invalid {
				fmt.Println(errorStatement)
				continue
			}
		} else {
			fmt.Println(errorStatement)
			continue
		}

		switch action {
		case "newanimal":
			{
				//word[1] => Name of animal word[2] = >type
				if !stringInSlice(words[2], getSupportedAnimalType()) {
					fmt.Println(errorStatement)
					continue
				}
				//"cow", "bird", "snake"
				switch words[2] {
				case "cow":
					userAnimalCollection[words[1]] = Cow{}
				case "bird":
					userAnimalCollection[words[1]] = Bird{}
				case "snake":
					userAnimalCollection[words[1]] = Snake{}

				}
				fmt.Println("Created or replaced it!")

			}
		case "query":
			{
				//word[1] => name of animal , word [2] = action
				if !stringInSlice(words[2], getSupportedActivity()) {
					fmt.Println(errorStatement)
					continue
				}
				//"eat", "move", "speak"
				if _, found := userAnimalCollection[words[1]]; found {

					switch words[2] {
					case "eat":
						userAnimalCollection[words[1]].Eat()
					case "move":
						userAnimalCollection[words[1]].Move()
					case "speak":
						userAnimalCollection[words[1]].Speak()

					}
				} else {
					fmt.Println(errorNotFound)
					continue
				}

			}
		}
	}
}
