package sort

import (
	"algorithms-data-structures/data-structures/queue"
	"fmt"
)

func MergeSort(slice []int) []int {
	divided := queue.New[[]int](len(slice))
	toDivide := queue.New[[]int](len(slice))
	toDivide.Enqueue(slice)

	for !toDivide.IsEmpty() {
		left, right := divide(toDivide.Dequeue())
		shouldDivide(left, toDivide, divided)
		shouldDivide(right, toDivide, divided)
	}

	//merged := stack.New[[]int](divided.Length())

	var result []int
	for !divided.IsEmpty() {
		first := divided.Dequeue()
		if len(first) == len(slice) {
			result = first
			break
		}

		// Check that there is still a sub-array to merge with the first.
		if !divided.IsEmpty() {
			second := divided.Dequeue()
			divided.Enqueue(merge(first, second))
		} else {
			divided.Enqueue(first)
		}
	}

	return result
}

func merge(sliceOne []int, sliceTwo []int) []int {
	merged := []int{}

	oneIndex, twoIndex := 0, 0
	for oneIndex < len(sliceOne) && twoIndex < len(sliceTwo) {
		if sliceOne[oneIndex] < sliceTwo[twoIndex] {
			merged = append(merged, sliceOne[oneIndex])
			oneIndex++
		} else {
			merged = append(merged, sliceTwo[twoIndex])
			twoIndex++
		}
	}

	// Add remaining values from each slice to merged slice as one slice was fully traversed and appended
	// to the slice containing the sorted, merged values.
	for oneIndex < len(sliceOne) {
		merged = append(merged, sliceOne[oneIndex])
		oneIndex++
	}

	for twoIndex < len(sliceTwo) {
		merged = append(merged, sliceTwo[twoIndex])
		twoIndex++
	}

	fmt.Printf("oneSlice: %s twoSlice: %s\n", fmt.Sprint(sliceOne), fmt.Sprint(sliceTwo))
	fmt.Printf("oneIndex: %d oneLength: %d\n", oneIndex, len(sliceOne))
	fmt.Printf("twoIndex: %d twoLength: %d\n", twoIndex, len(sliceTwo))
	fmt.Printf("merged: %s\n---\n", fmt.Sprint(merged))

	return merged
}

// divide is a helper function that will take in a single slice and divide it in half and return the separate halves.
// if the provided slice is empty, it will return nil for both returned halves.
func divide(slice []int) ([]int, []int) {
	high := len(slice)
	if high < 1 { // if empty slice, unable to divide
		return nil, nil
	}

	middle := high / 2

	return slice[:middle], slice[middle:]
}

// shouldDivide determines if the provided slice should be further divided, if necessary, it will be added
// to the toDivide queue for further division. If fully divided, it will be added ot the divided queue.
func shouldDivide(slice []int, toDivide *queue.Queue[[]int], divided *queue.Queue[[]int]) {
	if len(slice) > 1 {
		toDivide.Enqueue(slice)
	} else if len(slice) == 1 {
		divided.Enqueue(slice)
	}
}
