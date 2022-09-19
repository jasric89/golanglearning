package main

import "fmt"

//* Create a structure to store items and their security tag state

type items struct {
	name          string
	securityState bool
}

//  - Security tags have two states: active (true) and inactive (false)

const (
	active   = true
	inactive = false
)

//* Create functions to activate and deactivate security tags using pointers

func securityTag(status *items) {
	securityStatus := &status.securityState
	switch *securityStatus {
	case active:
		*securityStatus = inactive
	case inactive:
		*securityStatus = active
	}
	fmt.Println(status)
}

//  - Call the checkout() function to deactivate all tags

func checkout(state *[]items) array { // *items will access the value of the items in the items list.
	for _, securityTag := range *state {
		securityTag.securityState = false // we want to access the date in the struct so we put a . in to go to the actual data without the dot were just saying to Go get to the struct.
		return state
		fmt.Println("state", securityTag.securityState)
	}
}

func main() {

	shoppingList := []items{
		{"eggs", active},
		{"bolts", inactive},
		{"taps", active},
		{"almonds", inactive},
	}

	//  - Print out the array/slice after each change

	fmt.Println(shoppingList)

	//  - Deactivate any one security tag in the array/slice

	securityTag(&shoppingList[0])

	//  - Call the checkout() function to deactivate all tags

	checkout(&shoppingList) //& will access the values of the Items list, without the & the compiler just gets a bunch of memory addresses.

}
