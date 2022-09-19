//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func assemblyLine(Parts string, slice []string ){
	fmt.Println()
	fmt.Println("----", Parts, "----")
	for i := 0; i < len(slice); i++ {
		items := slice[i]
		fmt.Println(items)
	 }
}

func printAssemblyline( Title string, parts []string) {
	fmt.Println()
	fmt.Println("---", Title, "---")
	for _, p:= range parts {
		fmt.Println(p) 
	}
}

func main() {

	parts := []string{"Screwdriver","Spanner","Hammer","Level"}
    assemblyLine("First Run", parts)

	parts = append(parts, "Jigsaw", "Saw")
    assemblyLine("Second Run", parts)
    
	
	printAssemblyline("All Parts", parts)
		
	parts = parts[4:]
	assemblyLine("Third Run", parts)


}
