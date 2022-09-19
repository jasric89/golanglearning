//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//func setShoppingList(string) string {}

func main() {
	var shoppinglist [5]string
	shoppinglist[0] = "Sausages"
	shoppinglist[1] = "Bannana"
	shoppinglist[2] = "Grapes"
	shoppinglist[3] = "Apple"
	shoppinglist[4] = "Pear"
	fmt.Println("The whole shopping list is", len(shoppinglist))
	fmt.Println("All items on the shopping list are", shoppinglist)
	fmt.Println("The last item on the shopping list is", shoppinglist[4])
}
