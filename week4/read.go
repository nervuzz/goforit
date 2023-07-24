/* Write a program which reads information from a file and represents it in
a slice of structs. Assume that there is a text file which contains a series
of names. Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a 'name' struct which has two fields, 'fname' for the first name,
and 'lname' for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will
successively read each line of the text file and create a struct which contains the first
and last names found in the file. Each struct created will be added to a slice, and after
all lines have been read from the file, your program will have a slice containing one
struct for each line in the file. After reading all lines from the file, your program should
iterate through your slice of structs and print the first and last names found in each struct.*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return input
}

// Read file content and split into lines
func readLinesFromFile(fname string) []string {
	const W_LINE_SEPARATOR string = "\r\n"
	const U_LINE_SEPARATOR string = "\n"

	barr, err := os.ReadFile(fname)

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(-1)
	}

	return strings.Split(
		// Cover both Windows&Unix line-breaks
		strings.ReplaceAll(string(barr), W_LINE_SEPARATOR, U_LINE_SEPARATOR),
		U_LINE_SEPARATOR)
}

// Declare 'name' struct
type name struct {
	fname, lname [20]string
}

// Create new struct of type 'name' basing on space-separated string
func toNameStruct(str string) name {
	separatedStr := strings.Fields(str)
	firstName := [20]string{separatedStr[0]}
	lastName := [20]string{separatedStr[1]}
	return name{firstName, lastName}
}

// Convert fixed-size array to string
func arrToStr(arr []string) string {
	return strings.Trim(strings.Join(arr, ""), "[ ]")
}

// String representation of 'name' struct
func nameStructItemToStr(n name) (string, string) {
	firstName := arrToStr(n.fname[:])
	lastName := arrToStr(n.lname[:])
	return firstName, lastName
}

func main() {
	seriesOfNames := make([]name, 0)

	fname := GetUserInput("Type the file name and hit [Enter]: ")
	lines := readLinesFromFile(fname)

	for _, item := range lines {
		seriesOfNames = append(seriesOfNames, toNameStruct(item))
	}

	fmt.Println("\nFirst and last names found in each struct:")
	for _, item := range seriesOfNames {
		fmt.Println(nameStructItemToStr(item))
	}
}
