package main

import "fmt"

func runningSum(nums [4]int){
 var a, b, c, d []int
	for i := 0; i < len(nums); i++ {
	a = nums [0] + 1;
	b = nums[1] + 2;
    c = nums[2] + 3;
	d = nums[3] + 4;
 } 
 fmt.Println(a, b, c, d)
 //return rob 
}

func main() {
	var nums [4]int    
    nums[0] = 1
    nums[1] = 2
    nums[2] = 3
    nums[3] = 4
    //X := runningSum(nums)
	//fmt.Println(X)
}