package main

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	current, global := 0, 0
	for i := 0; i < len(gas); i++ {
		current += gas[i] - cost[i]
		global += gas[i] - cost[i]
		if current < 0 {
			current = 0
			start = i + 1
		}
	}

	if global < 0 {
		return -1
	}
	return start
}
