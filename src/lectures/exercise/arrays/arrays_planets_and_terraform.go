package main 

import "fmt"
// Important to note this program is designed to error in Lesson 17 and 26 is when this gets corrected with Slices and Pointers. Will come back to this. 
// Terraform accompolishes Nothing
func terraform(planets [8]string){
  for i := range planets {
	planets[i] := "New" + planets[i] {
	}
  }
}

func main () {
 planets :=[...]string {
	"Mercury",
	"Venus",
	"Earth",
	"Mars",
	"Jupiter",
	"Saturn",
	"Uranus",
	"Neptune",
 }
 terraform(planets)
 fmt.Println(planets)
}