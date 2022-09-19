package main

/*
Package Strings converting string cases.
The strings package provides the function for changing cases of strings.

ToLower(str) This function returns a new string containing the characters in the specified string mapped to lowercase
ToUpper(str) This function returns a new string containing the characters in the specified string mapped to uppercase
Title(str) This function converts the specific string so that the first character of each word is uppercase and the remaining characters are lowercase.
ToTitle(str) This function returns a new string containing the characters in the specified string mapped to title case
*/
import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for sailing"

	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description)) // Converts each character of the first word to upper case.
	/*
		The bellow code converts the characters to their unicode specification, and outputs those numbers depending on lowercase or uppercase.
	*/
	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
}
