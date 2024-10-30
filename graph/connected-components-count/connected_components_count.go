package connected_components_count

import (
	"github.com/sosalejandro/algo-practice/data-structures/common"
	"github.com/sosalejandro/algo-practice/graph"
)

// IntItem represents an item in the stack or queue.
type IntItem struct {
	value int
}

// NewIntItem creates a new IntItem
func NewIntItem(value int) *IntItem {
	return &IntItem{value: value}
}

// Equals checks if two IntItems have the same value.
func (i *IntItem) Equals(item common.Item[int]) bool {
	return i.value == item.Value()
}

// Value returns the value of the IntItem.
func (i *IntItem) Value() int {
	return i.value
}

// IsEmpty checks if the IntItem is empty.
func (i *IntItem) IsEmpty() bool {
	return i.value == 0
}

// ItemFactory function for IntItem.
var IntItemFactory common.ItemFactory[int] = func(value int) common.Item[int] {
	return NewIntItem(value)
}

// ConnectedComponentsCount returns the number of connected components in a graph.
// The graph (g) is represented as an adjacency list.
// The itemFactory parameter is used to create items for the transporter.
// The function returns an error if the transporter encounters an error.
func ConnectedComponentsCount[T comparable](strategy graph.TraversalStrategy, g map[T][]T, itemFactory common.ItemFactory[T]) (count int, err error) {
	if len(g) == 0 {
		return 0, nil
	}

	transporter := graph.NewTransporter[T](strategy, itemFactory)
	if transporter == nil {
		return 0, nil
	}

	visited := make(map[T]bool)

	for node := range g {
		if !visited[node] {
			count++
			transporter.Add(node)
			visited[node] = true

			for !transporter.IsEmpty() {
				currentItem, err := transporter.Next()
				if err != nil {
					return count, err
				}
				current := currentItem.Value()

				for _, neighbor := range g[current] {
					if !visited[neighbor] {
						visited[neighbor] = true
						transporter.Add(neighbor)
					}
				}
			}
		}
	}

	return count, nil
}
