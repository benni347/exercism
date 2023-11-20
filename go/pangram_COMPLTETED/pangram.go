package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	m := make(map[rune]int)
	for _, c := range input {
		if c >= 'a' && c <= 'z' {
			m[c]++
		}
	}
	return len(m) == 26
}
