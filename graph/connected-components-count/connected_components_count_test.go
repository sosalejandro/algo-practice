package connected_components_count

import (
	"testing"

	"github.com/sosalejandro/algo-practice/graph"
)

func TestConnectedComponentsCountDFS(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[int][]int
		expected int
	}{
		{
			name:     "Empty Graph",
			graph:    map[int][]int{},
			expected: 0,
		},
		{
			name: "Single Node Graph",
			graph: map[int][]int{
				1: {},
			},
			expected: 1,
		},
		{
			name: "Disconnected Graph",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
			},
			expected: 2,
		},
		{
			name: "Connected Graph",
			graph: map[int][]int{
				1: {3},
				2: {4},
				3: {2},
				4: {1},
			},
			expected: 1,
		},
		{
			name: "Graph with Cycles",
			graph: map[int][]int{
				1: {2},
				2: {3},
				3: {1},
			},
			expected: 1,
		},
		{
			name: "Graph with Multiple Components",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
				5: {},
			},
			expected: 3,
		},
		{
			name: "Two disconnected nodes",
			graph: map[int][]int{
				1: {},
				2: {},
			},
			expected: 2,
		},
		{
			name: "Graph with Multiple Components 2",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {},
				4: {6},
				5: {6},
				6: {4, 5, 7, 8},
				7: {6},
				8: {6},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_StackTraversal", func(t *testing.T) {
			result, err := ConnectedComponentsCount(graph.StackTraversal, tt.graph, IntItemFactory)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestConnectedComponentsCountBFS(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[int][]int
		expected int
	}{
		{
			name:     "Empty Graph",
			graph:    map[int][]int{},
			expected: 0,
		},
		{
			name: "Single Node Graph",
			graph: map[int][]int{
				1: {},
			},
			expected: 1,
		},
		{
			name: "Disconnected Graph",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
			},
			expected: 2,
		},
		{
			name: "Connected Graph",
			graph: map[int][]int{
				1: {2},
				2: {3},
				3: {4},
				4: {1},
			},
			expected: 1,
		},
		{
			name: "Graph with Cycles",
			graph: map[int][]int{
				1: {2},
				2: {3},
				3: {1},
			},
			expected: 1,
		},
		{
			name: "Graph with Multiple Components",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
				5: {},
			},
			expected: 3,
		},
		{
			name: "Two disconnected nodes",
			graph: map[int][]int{
				1: {},
				2: {},
			},
			expected: 2,
		},
		{
			name: "Graph with Multiple Components 2",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {},
				4: {6},
				5: {6},
				6: {4, 5, 7, 8},
				7: {6},
				8: {6},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_QueueTraversal", func(t *testing.T) {
			result, err := ConnectedComponentsCount(graph.QueueTraversal, tt.graph, IntItemFactory)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
