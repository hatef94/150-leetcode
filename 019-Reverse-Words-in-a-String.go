package main

import (
	"strings"
)

func main() {
	reverseWords("  hello    world  ")
}
func reverseWords(s string) string {
	charArr := []byte(strings.TrimSpace(s))
	slow, fast := 0, 0
	for ; fast < len(charArr); fast++ {
		if charArr[fast] != ' ' {
			charArr[slow] = charArr[fast]
			slow++
		} else {
			if fast < len(charArr)-1 && charArr[fast+1] != ' ' {
				charArr[slow] = charArr[fast]
				slow++
			}
		}
	}
	reverse(charArr, 0, slow-1)
	i, j := 0, 0
	for j < slow {
		if charArr[j] != ' ' {
			j++
		} else {
			reverse(charArr, i, j-1)
			j++
			i = j
		}
	}
	reverse(charArr, i, j-1)
	return string(charArr[0:slow])
}

func reverse(arr []byte, i, j int) {
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
}
