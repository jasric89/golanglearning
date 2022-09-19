//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

type rectangle struct {
	lengh  float64
	bredth float64
	height float64
	colour string
}

func area(bredth float64, lengh float64) float64 {
	return lengh * bredth
}
func sum(bredth float64, lengh float64, height float64) float64 {
	return bredth + (lengh * 2) + (height * 2)
}
func double(x float64) float64 {
	return x + x
}
func main() {
	var (
		rect = rectangle{lengh: 10.5, bredth: 25.10, height: 10.5}
	)
	fmt.Println(rectangle{10.5, 25.10, 10.5, "red"})
	fmt.Println("The Area of the rectangle is", area(rect.lengh, rect.bredth))
	fmt.Println("The Peremeter of a rectangle is", sum(rect.bredth, rect.height, rect.lengh))
	rectangledoubled := double(area(rect.lengh, rect.bredth) + sum(rect.bredth, rect.height, rect.lengh))
	fmt.Println("The rectangle doubled is", rectangledoubled)
}
