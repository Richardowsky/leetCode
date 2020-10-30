package longestSubstring

import "math"

func lengthOfLongestSubstring(s string) int {
	firstCharacter := 0
	nextCharacter := 0
	longest := 0
	set := make(map[string]bool)
	for firstCharacter < len(s) {
		_, ok := set[string(s[firstCharacter])]
		if ok {
			for s[nextCharacter] != s[firstCharacter] {
				delete(set, string(s[nextCharacter]))
				nextCharacter++
			}
			nextCharacter++
		} else {
			set[string(s[firstCharacter])] = true
			longest = int(math.Max(float64(longest), float64(firstCharacter-nextCharacter+1)))
		}
		firstCharacter++
	}
	return longest
}
