package main

import "fmt"

func main() {
	x := productExceptSelf([]int{1, 2})
	fmt.Println(x)
}

func productExceptSelf(nums []int) []int {
	answers := make([]int, len(nums), len(nums))
	answers[len(answers)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		answers[i] = nums[i+1] * answers[i+1]
	}

	sofar := 1
	for i := 0; i < len(answers); i++ {
		answers[i] = sofar * answers[i]
		sofar *= nums[i]
	}
	return answers
}
