package dfs

import (
	simple_stack "github.com/sosalejandro/algo-practice/data-structures/simple-stack"
	"github.com/sosalejandro/data-structures/common"
)

// StackStringItem represents an item in the stack.
type StackStringItem struct {
	value string
}

// NewStackStringItem creates a new StackStringItem.
func NewStackStringItem(value string) *StackStringItem {
	return &StackStringItem{value: value}
}

// Equals checks if two stackStringItems have the same value.
func (ssi *StackStringItem) Equals(other common.Item[string]) bool {
	if other == nil {
		return false
	}
	return ssi.value == other.Value()
}

// Value returns the value of the StackStringItem.
func (ssi *StackStringItem) Value() string {
	return ssi.value
}

// IsEmpty checks if the StackStringItem is empty.
func (ssi *StackStringItem) IsEmpty() bool {
	return ssi.value == ""
}

// ItemFactory function for stackStringItem.
var StackStringItemFactory common.ItemFactory[string] = func(value string) common.Item[string] {
	return NewStackStringItem(value)
}

// DeepFirstSearch is a generic DFS algorithm using the Item and ItemFactory function type.
func DeepFirstSearch[T comparable](graph map[T][]T, start T, itemFactory common.ItemFactory[T]) []T {
	if itemFactory(start).IsEmpty() || graph[start] == nil {
		return []T{}
	}

	visited := make(map[T]bool)
	response := make([]T, 0)

	stack := simple_stack.NewStack[T, common.Item[T]]()
	stack.Push(itemFactory(start))

	for !stack.IsEmpty() {
		currentItem, err := stack.Pop()
		if err != nil {
			break
		}
		current := currentItem.Value()

		if _, exists := graph[current]; !exists {
			continue
		}
		if _, ok := visited[current]; ok {
			continue
		}

		visited[current] = true
		response = append(response, current)

		for _, node := range graph[current] {
			if !visited[node] {
				stack.Push(itemFactory(node))
			}
		}
	}
	return response
}
