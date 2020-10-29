package addTwoNumbers

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var list = &ListNode{2, &ListNode{4, &ListNode{3, &ListNode{}}}}
var list2 = &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{}}}}

func TestList(t *testing.T) {
	excpect := &ListNode{7, &ListNode{0, &ListNode{8, &ListNode{}}}}
	assert.Equal(t, addTwoNumbers(list, list2), excpect)
}

func Benchmark(b *testing.B) {
	benchmarkZero(list, list2, b)
}

func benchmarkZero(l1 *ListNode, l2 *ListNode, b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
