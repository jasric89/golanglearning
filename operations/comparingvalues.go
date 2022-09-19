/* This program is comparing the values of variable first and the variable second. As seeing if they are equal not equal to, less than or more than the value.
Then is is outputting true or false to the result.
Also the const variable has a . in but is listed as an int. This is because go is disregarding the .00 if it was .01 then the program would get a compile error.
*/

package main

import "fmt"

func main() {
	first := 100          //Short hand variable decleration
	const second = 200.00 // creating a constant value variable you do not use : when assigning a value to a constant.

	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
}
