package sort

import (
	"algorithms-data-structures/data-structures/queue"
	"fmt"
)

// QuickSort performs the quick sort algorithm using the last element as the pivot
// on a provided slice and returns the sorted slice upon completion.
// if a nil slice is provided, an empty slice of ints will be returned.
func QuickSort(slice []int) []int {
	if slice != nil {
		if len(slice) == 1 {
			return slice
		}

		lowHighQueue := queue.Queue[int]{}
		lowHighQueue.EnqueueAll(0, len(slice)-1) // enqueue starting low and high values
		//lowHighQueue.EnqueueAll(4, 6)
		for !lowHighQueue.IsEmpty() {
			low, high := lowHighQueue.Dequeue(), lowHighQueue.Dequeue()
			left, right := partition(&slice, low, high)

			if shouldPartition(left) {
				lowHighQueue.EnqueueAll(left[0], left[1])
			}
			if shouldPartition(right) {
				lowHighQueue.EnqueueAll(right[0], right[1])
			}
		}

		return slice
	}

	return []int{}
}

// partition is a helper function to perform the bulk of the quick sort algorithm and solve for the
// next set of partitions being those below and above the pivot.
// the format of the return value is [low, high], [low, high] with the first low/high being the partition
// that is smaller than that of the pivot and the other being the values above the pivot.
func partition(slice *[]int, low int, high int) ([2]int, [2]int) {
	fmt.Printf("--- Partition - Low: %d High: %d ---\n", low, high)
	pivot, initialLow, initialHigh := high, low, high
	pivotValue := (*slice)[pivot]
	high = high - 1
	// low 0 high 2
	// 2 634 233
	for low < high {
		lowValue := (*slice)[low]
		highValue := (*slice)[high]
		if lowValue > pivotValue && highValue < pivotValue {
			fmt.Printf("Swapping %d with %d\n", lowValue, highValue)
			swap(slice, low, high)
		}

		if lowValue < pivotValue {
			low += 1
		}

		if highValue > pivotValue {
			high -= 1
		}
	}
	fmt.Printf("Before pivot swap: %s Low: %d High: %d\n", fmt.Sprint((*slice)), low, high)
	// do not swap pivot if all values were less than pivot
	if (*slice)[low] > pivotValue {
		swap(slice, low, pivot)
	} else {
		low += 1
	}

	fmt.Printf("After pivot swap: %s Low: %d High: %d\n", fmt.Sprint((*slice)), low, high)

	fmt.Printf("Left: %s Right: %s Pivot: %d\n", fmt.Sprint((*slice)[initialLow:low]), fmt.Sprint((*slice)[low+1:initialHigh+1]), pivotValue)

	leftPartition := [2]int{initialLow, low - 1}
	rightPartition := [2]int{low + 1, initialHigh}

	return leftPartition, rightPartition
}

// swap is a helper function to swap two values in place at the provided indexes with one anothers value.
func swap(slice *[]int, indexOne int, indexTwo int) {
	oldLast := (*slice)[indexOne]
	(*slice)[indexOne] = (*slice)[indexTwo]
	(*slice)[indexTwo] = oldLast
}

// shouldPartition given an array containing a low and high value for a possible partition, determines if we should
// perform the partitioning algorithm on said partition.
func shouldPartition(arr [2]int) bool {
	if intMathAbs(arr[0], arr[1]) > 1 {
		return true
	}

	return false
}

// intMathAbs is an implementation of math.Abs for int values.
func intMathAbs(numOne int, numTwo int) int {
	if numOne > numTwo {
		return numOne - numTwo
	}
	return numTwo - numOne
}
