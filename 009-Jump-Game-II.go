package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	reach, maxx := nums[0], nums[0]
	jumps := 0
	for i, num := range nums {
		if i > reach {
			jumps++
			reach = maxx
		}
		maxx = int(math.Max(float64(maxx), float64(i+num)))
	}
	return jumps + 1
}
