package reverseInteger

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var x = 654321

func TestList(t *testing.T) {
	excpect := 123456
	assert.Equal(t, reverse(x), excpect)
}

func Benchmark(b *testing.B) {
	benchmarkZero(x, b)
}

func benchmarkZero(x int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(x)
	}
}
