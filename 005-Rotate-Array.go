package main

func rotate(nums []int, k int) {
	i, j := 0, len(nums)-1
	k = k % len(nums)
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}

	i, j = 0, k-1
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}

	i, j = k, len(nums)-1
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}
