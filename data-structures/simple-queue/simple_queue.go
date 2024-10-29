package simple_queue

import (
	"errors"

	"github.com/sosalejandro/algo-practice/data-structures/common"
)

// QueueItem represents an item that can be stored in the queue.
// It requires methods to check equality and retrieve the item’s underlying value.
// type QueueItem[T any] interface {
// 	Equals(item common.Item[T]) bool
// 	Value() T
// 	IsEmpty() bool // To allow checking if the item itself is effectively empty
// }

// Queue represents a generic queue that holds elements of any type implementing QueueItem.
type Queue[T any, I common.Item[T]] struct {
	elements []I
}

// NewQueue creates a new instance of Queue.
func NewQueue[T any, I common.Item[T]]() *Queue[T, I] {
	return &Queue[T, I]{
		elements: make([]I, 0),
	}
}

// Enqueue adds an element to the end of the queue.
// It returns an error if the element is nil or empty.
func (q *Queue[T, I]) Enqueue(element I) error {
	if element.IsEmpty() {
		return errors.New("element cannot be nil or empty")
	}

	q.elements = append(q.elements, element)
	return nil
}

// Dequeue removes and returns the element at the front of the queue.
// It returns an error if the queue is empty.
func (q *Queue[T, I]) Dequeue() (I, error) {
	var zero I
	if len(q.elements) == 0 {
		return zero, errors.New("queue is empty")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

// Peek returns the element at the front of the queue without removing it.
func (q *Queue[T, I]) Peek() I {
	var zero I
	if len(q.elements) == 0 {
		return zero
	}

	return q.elements[0]
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T, I]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T, I]) Size() int {
	return len(q.elements)
}

// Clear removes all elements from the queue.
func (q *Queue[T, I]) Clear() {
	q.elements = make([]I, 0)
}

// ToSlice returns a slice containing all elements in the queue.
func (q *Queue[T, I]) ToSlice() []I {
	return q.elements
}

// Contains checks if the queue contains the specified element.
// It returns true if the element is found in the queue, and false otherwise.
func (q *Queue[T, I]) Contains(element I) bool {
	if element.IsEmpty() {
		return false
	}

	for _, curr := range q.elements {
		if curr.Equals(element) {
			return true
		}
	}
	return false
}
