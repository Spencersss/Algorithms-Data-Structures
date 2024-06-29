package bfs

import (
	"algorithms-data-structures/data-structures/graph"
	"algorithms-data-structures/data-structures/queue"
	"slices"
)

// BreadthFirstSearch performs the BFS algorithm against the provided graph and root node v.
// Once complete, returns the traversal order.
func BreadthFirstSearch(g graph.Graph, v int) []int {
	visited := []int{v}
	visiting := queue.Queue{}
	visiting.Enqueue(v)

	for !visiting.IsEmpty() {
		curNode := visiting.Dequeue()
		for neighbor, edge := range g.Edges[curNode] {
			// If neighbor not visited and edge is present making it a valid neighbor, proceed.
			if !slices.Contains(visited, neighbor) && edge == 1 {
				visiting.Enqueue(neighbor)
				visited = append(visited, neighbor)
			}
		}
	}

	return visited
}
