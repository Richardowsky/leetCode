package palindromeNumber

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var str = 10101

func TestList(t *testing.T) {
	excpect := true
	assert.Equal(t, isPalindrome(str), excpect)
}

func Benchmark(b *testing.B) {
	benchmarkZero(str, b)
}

func benchmarkZero(str int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(str)
	}
}
