package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueueWith[string](nil)

	assert.True(t, q.Empty())
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)

	assert.NotEmpty(t, q.Dequeue())
	assert.Empty(t, q.Dequeue())
}

func TestQueueString(t *testing.T) {
	expectation := []string{"h", "e", "l", "l", "o"}

	q := NewQueueWith(expectation)
	actual := make([]string, 0)
	for !q.Empty() {
		actual = append(actual, q.Dequeue())
	}

	assert.Equal(t, expectation, actual)
}

func TestQueueInt(t *testing.T) {
	expectation := []int{1, 2, 3, 4, 5}

	q := NewQueue[int]()
	for _, n := range expectation {
		q.Enqueue(n)
	}
	actual := make([]int, 0)
	for !q.Empty() {
		actual = append(actual, q.Dequeue())
	}

	assert.Equal(t, expectation, actual)
}
