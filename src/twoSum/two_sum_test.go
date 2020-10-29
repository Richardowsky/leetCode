package twoSum

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var nums = []int{2, 7, 11, 15}
var target = 9
var nums2 = []int{3, 2, 4}
var target2 = 6

func Test1(t *testing.T) {
	excpect := []int{1, 0}
	assert.Equal(t, twoSum(nums, target), excpect)
}

func Test2(t *testing.T) {
	excpect := []int{2, 1}
	assert.Equal(t, twoSum(nums2, target2), excpect)
}

func Benchmark1(b *testing.B) {
	benchmarkZero(target, nums, b)
}

func Benchmark2(b *testing.B) {
	benchmarkZero(target2, nums2, b)
}

func benchmarkZero(target int, nums []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
