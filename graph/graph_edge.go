package graph

// GraphType defines the type of graph to generate
type GraphType int

const (
	Directional GraphType = iota
	Bidirectional
)

// GenerateGraphFromEdges creates an adjacency list from an array of edges
// where for every edge in edges, edge[0] is the source and edge[1] is the destination.
// It is implicitly expected that every edge is an array with 2 elements; otherwise, it will be omitted.
func GenerateGraphFromEdges[T comparable](edges [][]T, graphType GraphType) map[T][]T {
	graph := make(map[T][]T)

	// Iterate over every edge entry
	for i := range edges {
		// Assuming every edge is a directional edge
		// We assume every entry is node A and node B
		// Node A is the source and node B is the destination

		if len(edges[i]) != 2 {
			continue
		}

		src, dst := edges[i][0], edges[i][1]

		// Create destination node if the entry doesn't yet exist
		if _, existsDstNode := graph[dst]; !existsDstNode {
			graph[dst] = make([]T, 0)
		}

		node, existsNode := graph[src]

		// Create the array if the entry doesn't yet exist
		if !existsNode {
			graph[src] = make([]T, 0)
		}

		// Add the entry to the graph
		graph[src] = append(node, dst)

		// If the graph is bidirectional, add the reverse edge as well
		if graphType == Bidirectional {
			nodeDst, existsNodeDst := graph[dst]
			if !existsNodeDst {
				graph[dst] = make([]T, 0)
			}
			graph[dst] = append(nodeDst, src)
		}
	}

	return graph
}
