package simple_stack

import (
	"errors"

	"github.com/sosalejandro/algo-practice/data-structures/common"
)

// StackItem represents an item that can be stored in the stack.
// It requires methods to check equality and retrieve the item’s underlying value.
// type common.Item[T any] interface {
// 	Equals(item common.Item[T]) bool
// 	Value() T
// 	IsEmpty() bool // To allow checking if the item itself is effectively empty
// }

// Stack represents a generic stack that holds elements of any type implementing StackItem.
type Stack[T any, I common.Item[T]] struct {
	elements []I
}

// NewStack creates a new instance of Stack.
func NewStack[T any, I common.Item[T]]() *Stack[T, I] {
	return &Stack[T, I]{
		elements: make([]I, 0),
	}
}

// Push adds an element to the top of the stack.
// It returns an error if the element is nil or empty.
func (s *Stack[T, I]) Push(element I) error {
	if element.IsEmpty() {
		return errors.New("element cannot be nil or empty")
	}

	s.elements = append(s.elements, element)
	return nil
}

// Pop removes and returns the element at the top of the stack.
// It returns an error if the stack is empty.
func (s *Stack[T, I]) Pop() (I, error) {
	var zero I
	if len(s.elements) == 0 {
		return zero, errors.New("stack is empty")
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

// Peek returns the element at the top of the stack without removing it.
func (s *Stack[T, I]) Peek() I {
	var zero I
	if len(s.elements) == 0 {
		return zero
	}

	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T, I]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T, I]) Size() int {
	return len(s.elements)
}

// Clear removes all elements from the stack.
func (s *Stack[T, I]) Clear() {
	s.elements = make([]I, 0)
}

// ToSlice returns a slice containing all elements in the stack.
func (s *Stack[T, I]) ToSlice() []I {
	return s.elements
}

// Contains checks if the stack contains the specified element.
// It returns true if the element is found in the stack, and false otherwise.
func (s *Stack[T, I]) Contains(element I) bool {
	if element.IsEmpty() {
		return false
	}

	for _, curr := range s.elements {
		if curr.Equals(element) {
			return true
		}
	}
	return false
}
