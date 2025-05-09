package main

import "math"

func canJump(nums []int) bool {
	reach := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > reach {
			return false
		}
		reach = int(math.Max(float64(reach), float64(i+nums[i])))
	}
	return true
}
