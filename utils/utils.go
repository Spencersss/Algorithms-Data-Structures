package utils

// New2DSlice creates a new 2D slice of size n x n.
func New2DSlice(n int) [][]int {
	newSlice := make([][]int, n)

	for newRow := 0; newRow < n; newRow++ {
		newSlice[newRow] = make([]int, n)
	}

	return newSlice
}
