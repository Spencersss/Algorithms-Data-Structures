package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessQuickSort(t *testing.T) {
	unsorted := []int{6000, 634, 7123, 233, 2, 920, 10000}
	unsortedTwo := []int{5000, 4000, 4500, 300, 450}

	result := QuickSort(unsorted)
	expected := []int{2, 233, 634, 920, 6000, 7123, 10000}

	resultTwo := QuickSort(unsortedTwo)
	expectedTwo := []int{300, 450, 4000, 5000, 4500}

	assert.NotNil(t, result, "QuickSort returned an unexpected nil slice.")
	assert.Equalf(t, expected, result, "Expected %s but found %s.", fmt.Sprint(expected), fmt.Sprint(result))

	assert.NotNil(t, resultTwo, "QuickSort returned an unexpected nil slice.")
	assert.Equalf(t, expectedTwo, resultTwo, "Expected %s but found %s.", fmt.Sprint(expectedTwo), fmt.Sprint(resultTwo))
}
