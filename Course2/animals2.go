package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		if scanner.Scan() {
			AnimalActivity := scanner.Text()
			Lst := strings.Fields(AnimalActivity)
			AnimalName := strings.ToLower(Lst[0])
			activity := strings.ToLower(Lst[1])
			animal := createAnimal(AnimalName)
			fmt.Println(CallAnimalActivity(animal, activity))
		}

	}

}

func createAnimal(animal string) Animal {
	switch animal {
	case "cow":
		return Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		return Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		return Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		return Animal{}
	}
}

func CallAnimalActivity(animal Animal, activity string) string {
	switch activity {
	case "eat":
		return animal.Eat()
	case "move":
		return animal.Move()

	case "speak":
		return animal.Speak()
	default:
		return "None"

	}

}

func (animal Animal) Eat() string {
	return animal.food

}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.noise

}
