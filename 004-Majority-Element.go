package main

func majorityElement(nums []int) int {
	m := nums[0]
	k := 1
	for _, num := range nums[1:] {
		if num == m {
			k++
		} else {
			k--
		}

		if k == 0 {
			m = num
		}
	}

	return m
}
