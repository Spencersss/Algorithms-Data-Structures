package dfs

import (
	"algorithms-data-structures/data-structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessDepthFirstSearch(t *testing.T) {
	testGraph := graph.New(4)
	testGraph.AddEdges(
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
		[]int{2, 0},
		[]int{2, 3},
		[]int{3, 3},
	)

	result := DepthFirstSearch(*testGraph, 2)
	expected := []int{2, 0, 3, 1}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
