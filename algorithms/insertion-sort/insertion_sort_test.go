package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessInsertionSort(t *testing.T) {
	unsorted := []int{5, 150, 23, 76, 77}

	result := InsertionSort(unsorted)
	expected := []int{5, 23, 76, 77, 150}

	assert.NotNil(t, result, "InsertionSort returned an unexpected nil slice.")
	assert.Equalf(t, expected, result, "Expected %s but found %s.", fmt.Sprint(expected), fmt.Sprint(result))
}
