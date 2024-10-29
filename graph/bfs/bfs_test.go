package bfs_test

import (
	"testing"

	"github.com/sosalejandro/algo-practice/graph/bfs"
)

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[string][]string
		start    string
		expected []string
	}{
		{
			name:     "Empty Graph",
			graph:    map[string][]string{},
			start:    "A",
			expected: []string{},
		},
		{
			name: "Single Node Graph",
			graph: map[string][]string{
				"A": {},
			},
			start:    "A",
			expected: []string{"A"},
		},
		{
			name: "Disconnected Graph",
			graph: map[string][]string{
				"A": {"B"},
				"B": {},
				"C": {"D"},
				"D": {},
			},
			start:    "A",
			expected: []string{"A", "B"},
		},
		{
			name: "Graph with Cycles",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"A"},
			},
			start:    "A",
			expected: []string{"A", "B", "C"},
		},
		{
			name: "Non-Existent Start Node",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
			},
			start:    "X",
			expected: []string{},
		},
		{
			name: "General Case",
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {"D", "E"},
				"C": {"F"},
				"D": {},
				"E": {"F"},
				"F": {},
			},
			start:    "A",
			expected: []string{"A", "B", "C", "D", "E", "F"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itemFactory := bfs.QueueStringItemFactory
			result := bfs.BreadthFirstSearch(tt.graph, tt.start, itemFactory)
			if !containsAll(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func containsAll(result, expected []string) bool {
	if len(result) != len(expected) {
		return false
	}
	resultMap := make(map[string]bool)
	for _, v := range result {
		resultMap[v] = true
	}
	for _, v := range expected {
		if !resultMap[v] {
			return false
		}
	}
	return true
}
