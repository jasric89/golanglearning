//Given the below class:
//class Cat:
//    species = 'mammal'
//    def init(self, name, age):
//        self.name = name
//        self.age = age
//1 Instantiate the Cat object with 3 cats
//2 Create a function that finds the oldest cat
//3 Print out: "The oldest cat is x years old.". x will be the oldest cat age by using the function in #2

package main

import "fmt"

type Cat struct {
    Name string
    Age  int
}

func main() {
    catOne := Cat{
        Name: "Kitty",
        Age:  20,
    }
    catTwo := Cat{
        Name: "casey",
        Age:  20,
    }
    catThree := Cat{
        Name: "Lily",
        Age:  20,
    }
   
}