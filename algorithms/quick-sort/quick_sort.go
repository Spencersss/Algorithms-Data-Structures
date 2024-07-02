package sort

import (
	"algorithms-data-structures/data-structures/queue"
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
	pivot, initialLow, initialHigh := high, low, high
	pivotValue := (*slice)[pivot]
	high = high - 1
	// low 0 high 2
	// 2 634 233
	for low < high {
		lowValue := (*slice)[low]
		highValue := (*slice)[high]
		if lowValue > pivotValue && highValue < pivotValue {
			swap(slice, low, high)
		}

		if lowValue < pivotValue {
			low += 1
		}

		if highValue > pivotValue {
			high -= 1
		}
	}
	// swap if left bound ended up being greater than pivot value prior to pivot swap
	if (*slice)[low] > pivotValue {
		swap(slice, low, pivot)
	} else {
		low += 1
	}
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
