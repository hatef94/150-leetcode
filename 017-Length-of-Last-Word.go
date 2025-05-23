package main

func main() {
	lengthOfLastWord("   fly me   to   the moon  ")
}
func lengthOfLastWord(s string) int {
	l := 0
	lastNotSpaceChar := len(s) - 1
	for ; lastNotSpaceChar >= 0; lastNotSpaceChar-- {
		if s[lastNotSpaceChar] != ' ' {
			break
		}
	}

	for s[lastNotSpaceChar] != ' ' {
		l++
		lastNotSpaceChar--
	}
	return l
}
