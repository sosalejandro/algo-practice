package has_path

import (
	"testing"

	"github.com/sosalejandro/algo-practice/graph"
)

func TestHasPath(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[string][]string
		src      string
		dst      string
		expected bool
	}{
		{
			name:     "Empty Graph",
			graph:    map[string][]string{},
			src:      "A",
			dst:      "B",
			expected: false,
		},
		{
			name: "Single Node Graph",
			graph: map[string][]string{
				"A": {},
			},
			src:      "A",
			dst:      "A",
			expected: true,
		},
		{
			name: "Disconnected Graph",
			graph: map[string][]string{
				"A": {"B"},
				"B": {},
				"C": {"D"},
				"D": {},
			},
			src:      "A",
			dst:      "D",
			expected: false,
		},
		{
			name: "Graph with Path",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"D"},
				"D": {},
			},
			src:      "A",
			dst:      "D",
			expected: true,
		},
		{
			name: "Graph without Path",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"D"},
			},
			src:      "A",
			dst:      "E",
			expected: false,
		},
		{
			name: "Graph with Cycles",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"A"},
			},
			src:      "A",
			dst:      "C",
			expected: true,
		},
		{
			name: "Non-Existent Start Node",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
			},
			src:      "X",
			dst:      "C",
			expected: false,
		},
		{
			name: "Non-Existent Destination Node",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
			},
			src:      "A",
			dst:      "Y",
			expected: false,
		},
		{
			name: "Non-Existent Neighbor Node",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C", "X"}, // "X" does not exist in the graph
				"C": {"D"},
				"D": {},
			},
			src:      "A",
			dst:      "D",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_StackTraversal", func(t *testing.T) {
			result, err := HasPath(graph.StackTraversal, tt.graph, tt.src, tt.dst, StringItemFactory)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})

		t.Run(tt.name+"_QueueTraversal", func(t *testing.T) {
			result, err := HasPath(graph.QueueTraversal, tt.graph, tt.src, tt.dst, StringItemFactory)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
