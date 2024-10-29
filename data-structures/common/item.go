// common/item.go
package common

// Item represents an item that can be stored in any data structure (stack or queue).
// It requires methods to check equality, retrieve the item’s value, and check if the item is empty.
type Item[T any] interface {
	// Equals checks if two items have the same value.
	Equals(item Item[T]) bool
	// Value returns the value of the item.
	Value() T
	// IsEmpty checks if the item is empty.
	IsEmpty() bool
}

// ItemFactory is a function type that creates a new item with the specified value.
type ItemFactory[T any] func(value T) Item[T]
