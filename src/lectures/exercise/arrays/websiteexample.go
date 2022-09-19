package main 

import "fmt"

func runningSum(nums){
	nums := [4]int{1,2,3,4}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}	
}

func main(){
	var nums = [4]int{1,2,3,4}
	runningSum(nums)
}