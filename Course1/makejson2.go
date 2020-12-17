package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter your name:")
	in := bufio.NewReader(os.Stdin)
	name, _ := in.ReadString('\n')
	fmt.Println("Enter your address:")
	in2 := bufio.NewReader(os.Stdin)
	address, _ := in2.ReadString('\n')

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	b, _ := json.Marshal(m)
	_ = b
	fmt.Println(string(b))
}
