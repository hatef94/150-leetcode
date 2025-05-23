package main

func longestCommonPrefix(strs []string) string {
	minIdx := 0
	for idx, str := range strs {
		if len(str) < len(strs[minIdx]) {
			minIdx = idx
		}
	}

	end := 0
	for i := 0; i < len(strs[minIdx]); i++ {
		isCommonPrefix := true
		for _, str := range strs {
			if str[end] != strs[minIdx][end] {
				isCommonPrefix = false
				break
			}
		}
		if !isCommonPrefix {
			break
		}
		end++
	}
	return strs[minIdx][0:end]
}
