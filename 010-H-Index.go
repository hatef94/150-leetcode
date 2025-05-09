package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{0}))
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for idx, citation := range citations {
		if citation != 0 {
			papers := len(citations) - idx
			tempH := math.Min(float64(papers), float64(citation))
			h = int(math.Max(tempH, float64(h)))
		}
	}
	return h
}
