package longestSubstring

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var str = "abcaabbrr"

func TestList(t *testing.T) {
	excpect := 3
	assert.Equal(t, lengthOfLongestSubstring(str), excpect)
}

func Benchmark(b *testing.B) {
	benchmarkZero(str, b)
}

func benchmarkZero(str string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(str)
	}
}
