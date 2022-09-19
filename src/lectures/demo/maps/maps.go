package main

import "fmt"



func main() {
 
  shopping_list := make(map[string]int)
  shopping_list["eggs"] = 11
  shopping_list["milk"] = 1
  shopping_list["bread"] += 1
  
  shopping_list["eggs"] += 1
  fmt.Println(shopping_list)

  delete(shopping_list, "milk")
  fmt.Println("Milk Taken off", shopping_list)
  fmt.Println("need", shopping_list["eggs"], "eggs")
  
  ceral, found := shopping_list["ceral"]
  fmt.Println("Need Ceral?")
  if !found {
	fmt.Println("We have", ceral, "boxes of ceral")
  } else {
	fmt.Println("yes we do need ceral")
  }
  
  totalCount := 0
  for _, amount := range shopping_list {
    totalCount += amount 
  }
  fmt.Println("There are", totalCount, "items on the shopping list")
}
