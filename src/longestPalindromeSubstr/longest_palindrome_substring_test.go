package longestPalindromeSubstr

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var str = "babad"

func TestList(t *testing.T) {
	excpect := "aba"
	assert.Equal(t, longestPalindrome(str), excpect)
}

func Benchmark(b *testing.B) {
	benchmarkZero(str, b)
}

func benchmarkZero(str string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome(str)
	}
}
