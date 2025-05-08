package main

func removeDuplicatesII(nums []int) int {
	last := 0
	for _, val := range nums {
		if last == 0 || last == 1 || nums[last-2] != val {
			nums[last] = val
			last++
		}
	}
	return last
}
