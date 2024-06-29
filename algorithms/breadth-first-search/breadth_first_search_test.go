package bfs

import (
	"algorithms-data-structures/data-structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessBreadthFirstSearch(t *testing.T) {
	testGraph := graph.New(6)
	testGraph.AddEdges(
		[]int{1, 2},
		[]int{1, 5},
		[]int{2, 3},
		[]int{2, 4},
		[]int{4, 2},
		[]int{3, 4},
		[]int{4, 5},
	)

	result := BreadthFirstSearch(*testGraph, 1)
	expected := []int{1, 2, 5, 3, 4}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
