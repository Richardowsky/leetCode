package longestPalindromeSubstr

import "math"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := int(math.Max(float64(len1), float64(len2)))
		if maxLen > end-start {
			start = i - ((maxLen - 1) / 2)
			end = i + (maxLen / 2)
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left > -1 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
