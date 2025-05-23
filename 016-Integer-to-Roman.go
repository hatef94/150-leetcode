package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var rta = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

var mts = []int{1000, 500, 100, 50, 10, 5, 1}

func main() {
	fmt.Println(intToRoman(8742))
}

func intToRoman(num int) string {
	result := strings.Builder{}
	numStr := strconv.Itoa(num)
	for idx, digitChar := range numStr {
		if digitChar == '0' {
			continue
		}
		digit := int(digitChar - '0')
		order := int(math.Pow(10, float64(len(numStr)-idx-1)))
		n := digit * order
		switch digitChar {
		case '4', '9':
			result.Write([]byte(rta[order]))
			result.Write([]byte(rta[n+order]))
		default:
			for n > 0 {
				for _, mx := range mts {
					if n-mx < 0 {
						continue
					}
					h := n / mx
					result.WriteString(strings.Repeat(rta[mx], h))
					n -= mx * h
					break
				}
			}
		}
	}
	return result.String()
}
