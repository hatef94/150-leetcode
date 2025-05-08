package main

func removeDuplicates(nums []int) int {
	k, idx := 0, 0
	for idx < len(nums) {
		if nums[idx] != nums[k] {
			k++
			nums[k] = nums[idx]
		}
		idx++
	}
	return k + 1
}
