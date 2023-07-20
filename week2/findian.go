// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character ‘i’,
// ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise.
// The program should not be case-sensitive, so it does not matter
// if the characters are upper-case or lower-case.

package main

import (
	"fmt"
	"os"
	"strings"
)

func startswithCharI(aString string) bool {
	if strings.HasPrefix(strings.ToLower(aString), "i") {
		return true
	} else {
		return false
	}
}

func endswithCharN(aString string) bool {
	if strings.HasSuffix(strings.ToLower(aString), "n") {
		return true
	} else {
		return false
	}
}

func containsCharA(aString string) bool {
	if strings.Contains(strings.ToLower(aString), "a") {
		return true
	} else {
		return false
	}
}

func findian(aString string) string {
	// No leading/trailing whitespaces
	strTrim := strings.TrimSpace(aString)
	if len(strTrim) <= 2 {
		// fmt.Println("Not enough chars for 'findian'")
		return "Not Found!"
	} else {
		// Neither first nor last character
		strSliced := strTrim[1 : len(strTrim)-1]

		if startswithCharI(strTrim) &&
			endswithCharN(strTrim) &&
			containsCharA(strSliced) {
			return "Found!"
		} else {
			return "Not Found!"
		}
	}
}

func main() {
	var aString string
	fmt.Printf("Type some string and hit [Enter]: ")
	_, err := fmt.Scan(&aString)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	fmt.Println(findian(aString))
}
