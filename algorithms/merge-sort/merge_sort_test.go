package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessMergeSort(t *testing.T) {
	test := []int{100, 40, 50, 20, 250, 222}
	testTwo := []int{2000, 60, 999}

	result := MergeSort(test)
	expected := []int{20, 40, 50, 100, 222, 250}

	resultTwo := MergeSort(testTwo)
	expectedTwo := []int{60, 999, 2000}

	assert.Equalf(t, expected, result, "Expected %s but found %s.", fmt.Sprint(expected), fmt.Sprint(result))
	assert.Equalf(t, expectedTwo, resultTwo, "Expected %s but found %s.", fmt.Sprint(expectedTwo), fmt.Sprint(resultTwo))
}

func TestSuccessDivide(t *testing.T) {
	test := []int{20, 30, 40, 50}

	left, right := divide(test)
	expectedLeft := []int{20, 30}
	expectedRight := []int{40, 50}

	assert.Equalf(t, expectedLeft, left, "Expected %s but found %s.", fmt.Sprint(expectedLeft), fmt.Sprint(left))
	assert.Equalf(t, expectedRight, right, "Expected %s but found %s.", fmt.Sprint(expectedRight), fmt.Sprint(right))
}
