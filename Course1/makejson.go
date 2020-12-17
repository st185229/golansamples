/*
* Write a program which prompts the user to first enter a name, and then enter an address.
* Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
* Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
 */

/*
 3 points will be given if a program is written.
 2 points will be given if the program compiles correctly.
 5 points will be given if the program correctly prints a JSON object with keys ("name", "address")
 and they are associated with the name and address that was entered.
*/

// Assumption name and address could be multi worded

//Declare variable to hold input string
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	contactDetailsMap := make(map[string]string)

	//Request user to input
	fmt.Print("Enter your name, press return when done: ")
	//read name string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		contactDetailsMap["name"] = scanner.Text()
	}
	//read address string
	fmt.Print("Enter your address, press return when done: ")
	if scanner.Scan() {
		contactDetailsMap["address"] = scanner.Text()
	}
	//Convert to json
	//jsonout, err := json.Marshal(contactDetailsMap)
	jsonout, err := json.MarshalIndent(contactDetailsMap, "", " ")
	if err == nil {
		stringjson := string(jsonout)
		fmt.Println(stringjson)
	}
}
