package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// create an array of size 3 but the slice must point to the first index
	mapUser := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter User Name: ")
	scanner.Scan()
	userName := scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	address := scanner.Text()

	mapUser["name"] = userName
	mapUser["address"] = address

	jsonString, _ := json.Marshal(mapUser)

	fmt.Println(string(jsonString))
}
