package bfs

import (
	"github.com/sosalejandro/algo-practice/data-structures/common"
	simple_queue "github.com/sosalejandro/algo-practice/data-structures/simple-queue"
)

// QueueStringItem represents an item in the queue.
type QueueStringItem struct {
	value string
}

// NewQueueStringItem creates a new QueueStringItem.
func NewQueueStringItem(value string) *QueueStringItem {
	return &QueueStringItem{value: value}
}

// Equals checks if two queueStringItems have the same value.
func (ssi *QueueStringItem) Equals(other common.Item[string]) bool {
	if other == nil {
		return false
	}
	return ssi.value == other.Value()
}

// Value returns the value of the QueueStringItem.
func (ssi *QueueStringItem) Value() string {
	return ssi.value
}

// IsEmpty checks if the QueueStringItem is empty.
func (ssi *QueueStringItem) IsEmpty() bool {
	return ssi.value == ""
}

// ItemFactory function for queueStringItem.
var QueueStringItemFactory common.ItemFactory[string] = func(value string) common.Item[string] {
	return NewQueueStringItem(value)
}

// BreadthFirstSearch is a generic BFS algorithm using the Item and ItemFactory function type.
func BreadthFirstSearch[T comparable](graph map[T][]T, start T, itemFactory common.ItemFactory[T]) []T {
	if itemFactory(start).IsEmpty() || graph[start] == nil {
		return []T{}
	}

	visited := make(map[T]bool)
	response := make([]T, 0)

	queue := simple_queue.NewQueue[T, common.Item[T]]()
	queue.Enqueue(itemFactory(start))

	for !queue.IsEmpty() {
		currentItem, err := queue.Dequeue()
		if err != nil {
			break
		}
		current := currentItem.Value()

		if _, isVisited := visited[current]; isVisited {
			continue
		}

		visited[current] = true
		response = append(response, current)

		for _, node := range graph[current] {
			if !visited[node] {
				queue.Enqueue(itemFactory(node))
			}
		}
	}
	return response
}
