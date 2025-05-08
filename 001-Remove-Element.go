package main

func removeElement(nums []int, val int) int {
	k, i := 0, 0
	for i < len(nums) {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
		i++
	}
	return k
}
