package main

import "fmt"

func main() {
	x := romanToInt("MCMXCIV")
	fmt.Println(x)
}

var m = map[int32]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := 0
	last := 'M'
	for _, v := range s {
		if m[v] > m[last] {
			n -= 2 * m[last]
		}
		last = v
		n += m[v]
	}
	return n
}
