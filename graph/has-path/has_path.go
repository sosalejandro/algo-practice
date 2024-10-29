package has_path

import (
	"github.com/sosalejandro/algo-practice/data-structures/common"
	"github.com/sosalejandro/algo-practice/graph"
)

// StringItem represents an item in the stack or queue.
type StringItem struct {
	value string
}

// NewStringItem creates a new StringItem.
func NewStringItem(value string) *StringItem {
	return &StringItem{value: value}
}

// Equals checks if two StringItems have the same value.
func (s *StringItem) Equals(item common.Item[string]) bool {
	return s.value == item.Value()
}

// Value returns the value of the StringItem.
func (s *StringItem) Value() string {
	return s.value
}

func (ssi *StringItem) IsEmpty() bool {
	return ssi.value == ""
}

// NewItem creates a new StringItem.
func (ssi *StringItem) NewItem(value string) common.Item[string] {
	return NewStringItem(value)
}

// ItemFactory function for StringItem.
var StringItemFactory common.ItemFactory[string] = func(value string) common.Item[string] {
	return NewStringItem(value)
}

// HasPath is a generic pathfinding function using the Transporter interface for different traversal strategies.
// It returns true if there is a path from src to dst in the graph.
// The graph (g) is represented as an adjacency list.
// The itemFactory parameter is used to create items for the transporter.
// The function returns an error if the transporter encounters an error.
func HasPath[T comparable](strategy graph.TraversalStrategy, g map[T][]T, src, dst T, itemFactory common.ItemFactory[T]) (bool, error) {
	// Check if the source or destination node does not exist in the graph
	if _, srcExists := g[src]; !srcExists {
		return false, nil
	}
	if _, dstExists := g[dst]; !dstExists {
		return false, nil
	}

	if src == dst {
		return true, nil
	}

	transporter := graph.NewTransporter[T](strategy, itemFactory)
	if transporter == nil {
		return false, nil
	}

	visited := make(map[T]bool)
	visited[src] = true

	// Initialize traversal by adding the source node to the transporter
	if err := transporter.Add(src); err != nil {
		return false, err
	}

	for !transporter.IsEmpty() {
		currentItem, err := transporter.Next()
		if err != nil {
			return false, err
		}
		current := currentItem.Value()

		// Explore neighbors of the current node
		for _, neighbor := range g[current] {
			// // Check if the neighbor exists in the graph
			if _, neighborExists := g[neighbor]; !neighborExists {
				continue
			}

			if neighbor == dst {
				return true, nil
			}
			if !visited[neighbor] {
				visited[neighbor] = true
				if err := transporter.Add(neighbor); err != nil {
					return false, err
				}
			}
		}
	}
	return false, nil
}
