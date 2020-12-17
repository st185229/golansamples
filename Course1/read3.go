package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	FirstName string
	LastName  string
}

func main() {
	fmt.Print("Enter the file name : ")
	reader := bufio.NewReader(os.Stdin)

	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to read from stdin")
	}

	fileName = strings.TrimSuffix(fileName, "\n")

	fmt.Println(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to read file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var names []Name

	for scanner.Scan() {
		names = append(names, Name{FirstName: strings.Fields(scanner.Text())[0],
			LastName: strings.Fields(scanner.Text())[1]})
	}

	file.Close()

	// print the first and last names found in each struct.
	for _, name := range names {
		fmt.Println(name.FirstName + " " + name.LastName)
	}
}
