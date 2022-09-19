package main

import (
"fmt"
)

//func runningSum(nums [4]int) []int{
//	rob := make([]int, 4)
//	P := 0
//	for i := 0; i < len(nums); i++ {
//		fmt.Println(nums[i])
//		 P += nums[i]  //iteration of the array value 
//		rob[i] = P 
//    }
//	return rob
//}

func runningSum(nums []int) []int {
    rob := make([]int, len(nums))
	P := 0
    for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		 P += nums[i]  //iteration of the array value 
		rob[i] = P 
    }
	return rob
}

func main(){
 var nums [4]int    
    nums[0] = 1
    nums[1] = 2
    nums[2] = 3
    nums[3] = 4
    X := runningSum(nums)
	fmt.Println(X)
}