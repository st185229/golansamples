package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise string
}
type Bird struct {
	food, locomotion, noise string
}
type Snake struct {
	food, locomotion, noise string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}
func (a Cow) Move() {
	fmt.Println(a.locomotion)
}
func (a Cow) Speak() {
	fmt.Println(a.noise)
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}
func (a Bird) Move() {
	fmt.Println(a.locomotion)
}
func (a Bird) Speak() {
	fmt.Println(a.noise)
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}
func (a Snake) Move() {
	fmt.Println(a.locomotion)
}
func (a Snake) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := make(map[string]Animal)
	cow := Cow{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Bird{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Snake{food: "mice", locomotion: "slither", noise: "hsss"}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter a three part command A or B:")
		fmt.Println("A) [newanimal <animal name> <animal type (bird, cow, snake)>]")
		fmt.Println("B) [query <animal name> <action (eat, move, speak)])")
		fmt.Print(">")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
		inputAction := strings.Split(userInput, " ")[0]

		switch inputAction {
		case "newanimal":
			animalName := strings.Split(userInput, " ")[1]
			animalType := strings.Split(userInput, " ")[2]

			switch animalType {
			case "bird":
				animals[animalName] = bird
			case "cow":
				animals[animalName] = cow
			case "snake":
				animals[animalName] = snake
			default:
				fmt.Println("Invalid newanimal, try again\n")
				continue
			}

			fmt.Println("Created it!\n")
		case "query":
			animalName := strings.Split(userInput, " ")[1]
			animalAction := strings.Split(userInput, " ")[2]
			switch animalAction {
			case "eat":
				animals[animalName].Eat()
			case "move":
				animals[animalName].Move()
			case "speak":
				animals[animalName].Speak()
			default:
				fmt.Println("Invalid query, try again\n")
				continue
			}
		default:
			fmt.Println("Invalid action, try again\n")
			continue
		}
	}
}
