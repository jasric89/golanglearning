package main

/*
Count(s, sub) This function returns an int that reports how many times the specified substring is found in the string s

Index(s, sub) These functions returns the index of the first or last occurance of a specified substring string
LastIndex(s, sub) within the string s, or -1 if there is no occurence of the substring string

IndexAny(s, chars) These functions return the first or last occurance of a character in the specified substring
LastIndexAny(s, chars) with the string S or -1 if there is no occurance

IndexFunc(s, func) The IndexFunc and LastIndexFunc functions use a custom function to inspect strings, using custom functions as shown bellow.
LastIndexFunc(s, func)

*/

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"

	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
}
