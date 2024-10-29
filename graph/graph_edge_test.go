package graph_test

import (
	"reflect"
	"testing"

	"github.com/sosalejandro/algo-practice/graph"
)

func TestGenerateGraphFromEdges(t *testing.T) {
	tests := []struct {
		name      string
		graphType graph.GraphType
		edges     [][]string
		expected  map[string][]string
	}{
		{
			name:      "Directional: Empty Edges",
			edges:     [][]string{},
			graphType: graph.Directional,
			expected:  map[string][]string{},
		},
		{
			name:      "Bidirectional: Empty Edges",
			edges:     [][]string{},
			graphType: graph.Bidirectional,
			expected:  map[string][]string{},
		},
		{
			name:      "Directional: Single Edge",
			graphType: graph.Directional,
			edges: [][]string{
				{"A", "B"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {},
			},
		},
		{
			name:      "Bidirectional: Single Edge",
			graphType: graph.Bidirectional,
			edges: [][]string{
				{"A", "B"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {"A"},
			},
		},
		{
			name:      "Bidirectional: Multiple Edges",
			graphType: graph.Bidirectional,
			edges: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "D"},
				{"C", "D"},
			},
			expected: map[string][]string{
				"A": {"B", "C"},
				"B": {"A", "D"},
				"C": {"A", "D"},
				"D": {"B", "C"},
			},
		},
		{
			name:      "Directional: Multiple Edges",
			graphType: graph.Directional,
			edges: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "D"},
				{"C", "D"},
			},
			expected: map[string][]string{
				"A": {"B", "C"},
				"B": {"D"},
				"C": {"D"},
				"D": {},
			},
		},
		{
			name:      "Directional: Disconnected Nodes",
			graphType: graph.Directional,
			edges: [][]string{
				{"A", "B"},
				{"C", "D"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {},
				"C": {"D"},
				"D": {},
			},
		},
		{
			name:      "Bidirectional: Disconnected Nodes",
			graphType: graph.Bidirectional,
			edges: [][]string{
				{"A", "B"},
				{"C", "D"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {"A"},
				"C": {"D"},
				"D": {"C"},
			},
		},
		{
			name:      "Directional: Graph with Cycles",
			graphType: graph.Directional,
			edges: [][]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"A"},
			},
		},
		{
			name:      "Graph with Cycles",
			graphType: graph.Bidirectional,
			edges: [][]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			expected: map[string][]string{
				"A": {"B", "C"},
				"B": {"A", "C"},
				"C": {"B", "A"},
			},
		},
		{
			name:      "Directional: Invalid Edge",
			graphType: graph.Directional,
			edges: [][]string{
				{"A", "B"},
				{"B"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {},
			},
		},
		{
			name:      "Bidirectional: Invalid Edge",
			graphType: graph.Bidirectional,
			edges: [][]string{
				{"A", "B"},
				{"B"},
			},
			expected: map[string][]string{
				"A": {"B"},
				"B": {"A"},
			},
		},
		{
			name:      "Directional: Duplicate Edges",
			graphType: graph.Directional,
			edges: [][]string{
				{"A", "B"},
				{"A", "B"},
			},
			expected: map[string][]string{
				"A": {"B", "B"},
				"B": {},
			},
		},
		{
			name:      "Bidirectional: Duplicate Edges",
			graphType: graph.Bidirectional,
			edges: [][]string{
				{"A", "B"},
				{"A", "B"},
			},
			expected: map[string][]string{
				"A": {"B", "B"},
				"B": {"A", "A"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := graph.GenerateGraphFromEdges(tt.edges, tt.graphType)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
