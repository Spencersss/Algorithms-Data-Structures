package dfs

import (
	"algorithms-data-structures/data-structures/graph"
	"algorithms-data-structures/data-structures/stack"
	"slices"
)

func DepthFirstSearch(g graph.Graph, v int) []int {
	visited := []int{v}
	visiting := stack.New(10)
	visiting.Add(v)

	for !visiting.IsEmpty() {
		curNode := visiting.Pop()
		for neighbor, edge := range g.Edges[curNode] {
			// If neighbor not visited and edge is present making it a valid neighbor, proceed.
			if !slices.Contains(visited, neighbor) && edge == 1 {
				visiting.Add(neighbor)
				visited = append(visited, neighbor)
			}
		}
	}

	return visited
}
