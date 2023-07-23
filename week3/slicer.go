// Write a program which prompts the user to enter integers and stores the
// integers in a sorted slice. The program should be written as a loop.
// Before entering the loop, the program should create an empty integer slice
// of size (length) 3. During each pass through the loop, the program prompts
// the user to enter an integer to be added to the slice. The program adds the
// integer to the slice, sorts the slice, and prints the contents of the slice
// in sorted order. The slice must grow in size to accommodate any number of
// integers which the user decides to enter. The program should only quit
// (exiting the loop) when the user enters the character ‘X’ instead of an integer.
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getUserInput() string {
	var input string
	fmt.Printf("Type an integer and hit [Enter] or [X] to quit: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return input
}

func sortSliceAsc(aSlice []int) []int {
	sort.Slice(aSlice, func(i, j int) bool {
		return aSlice[i] < aSlice[j]
	})
	return aSlice
}

func main() {
	aSlice := make([]int, 0)

	for {
		userInput := getUserInput()
		aInt, err := strconv.Atoi(userInput)

		if err != nil {
			if userInput == "X" {
				fmt.Println("Exit on user request..")
			} else {
				fmt.Println("It's not an integer!")
				continue
			}
			break
		}
		aSlice = sortSliceAsc(append(aSlice, aInt))
		fmt.Println("Sorted slice: ", aSlice)
	}
}
