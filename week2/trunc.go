// Write a program which prompts the user to enter a floating point number and prints
// the integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.
package main

import "fmt"

func getUserInput() float64 {
	var aFloat float64
	fmt.Printf("Type a float and hit [Enter]: ")
	_, err := fmt.Scan(&aFloat)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return aFloat
}

func _main() {
	for i := 0; i < 2; i++ {
		fmt.Println("Truncated input: ", int(getUserInput()))
	}
}
