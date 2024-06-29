package graph

type Graph struct {
	Edges [][]int
}

func New(numNodes int) *Graph {
	newGraph := &Graph{
		Edges: make([][]int, numNodes),
	}

	for newRow := 0; newRow < numNodes; newRow++ {
		newGraph.Edges[newRow] = make([]int, numNodes)
	}

	return newGraph
}

func (g *Graph) AddEdge(n1 int, n2 int) {
	g.Edges[n1][n2] = 1
}

func (g *Graph) AddEdges(edges ...[]int) {
	for _, nodePair := range edges {
		if g.GetEdge(nodePair[0], nodePair[1]) == 0 {
			g.AddEdge(nodePair[0], nodePair[1])
		}
	}
}

func (g *Graph) RemoveEdge(n1 int, n2 int) {
	if g.GetEdge(n1, n2) == 1 {
		g.Edges[n1][n2] = 0
	}
}

// GetEdge determines if there is an edge that is directed from node n1 to node n2.
// If found returns 1, if not found returns 0.
func (g *Graph) GetEdge(n1 int, n2 int) int {
	return g.Edges[n1][n2]
}

// EdgeExists determines if there is an edge between vertex n1 and vertex n2 regardless of direction.
// If found returns true, if not returns false.
func (g *Graph) EdgeExists(n1 int, n2 int) bool {
	return g.Edges[n1][n2] == 1 || g.Edges[n2][n1] == 1
}
