// Write a program which prompts the user to first enter a name, and then
// enter an address. Your program should create a map and add the name
// and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map,
// and then your program should print the JSON object.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return input
}

func main() {
	aMap := make(map[string]string)

	aMap["name"] = getUserInput("Type the name and hit [Enter]: ")
	aMap["address"] = getUserInput("Type the address and hit [Enter]: ")

	aJson, err := json.Marshal(aMap)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(-1)
	}

	fmt.Printf("The JSON object: %s", aJson)
}
