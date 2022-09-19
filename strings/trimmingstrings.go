package main

/*
The process of trimming strings is to remove leading and trailing characters from a string and is most often used to remove whitespace

TrimSpace(s) This function returns the string s without leading or trailing whitespace characters

Trim(s, set) This function returns a string from which any leading or trailing characters contained in the string set are removed from the string s

TrimeLeft(s, set) This function returns the string s without any leading character contained in the string set. This function matches any of the specified characters-use
                 the TrimPrefix function to remove a complete substring

TrimRight(s, set) This function returns the string s without any leading character contained in the string set. This function matches any of the specified characters-use
                 the TrimPrefix function to remove a complete substring

TrimPrefix(s, prefix) This function returns the string s after removing the specified prefix string.
                      This function removes the complete prefix string - use the TrimLeft function to remove characters from a set

TrimSuffix(s, suffix) This function returns the string s after removing the specified suffix string
                      This function removes the complete prefix string - use the TrimRight function to remove characters from a set

TrimFunc(s, func) This function returns the string s from which any leading or trailing character for which a custom function returns true are removed

TrimLeftFunc(s, func) This function returns the string s from which any leading character for which a custom function returns true are removed

TrimRightFun(s, func) This function returns the string s from which any trailing character for which a custom function returns true are removed

*/

import (
	"fmt"
	"strings"
)

func main() {
	/*
		The user may have not realised they have pressed the spacebar when entering their name so this is what the TrimSpace func does takes the whitespace out making the username
		usable.
	*/
	username := " Alice"
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>"+trimmed+"<<")

	///

	/*
	   Here the following characters were specified, Asno so the Trim function will go through the string and match up those characters with what is specified and take them out.
	   Once it has got to the end of the string it will stop.
	*/
	description := "A boat for one person"
	trimmed2 := strings.Trim(description, "Asno")
	fmt.Println("Trimmed:", trimmed2)

	prefixTrimmed := strings.TrimPrefix(description, "A Boat")
	wrongPrefix := strings.TrimPrefix(description, "A Hat")

	fmt.Println("Trimmed:", prefixTrimmed)
	fmt.Println("Not Trimmed:", wrongPrefix)

}
