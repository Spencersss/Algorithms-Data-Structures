package sort

// InsertionSort performs the insertion sort algorithm on a provided slice and returns the sorted slice.
// if a nil slice is provided, an empty slice of ints will be returned.
func InsertionSort(slice []int) []int {
	if slice != nil {
		if len(slice) == 1 {
			return slice
		}

		// While curValue not sorted
		for curIndex := 1; curIndex < len(slice); curIndex++ {
			for tempCurIndex := curIndex; slice[tempCurIndex] < slice[tempCurIndex-1]; tempCurIndex-- {
				oldLast := slice[tempCurIndex-1]
				slice[tempCurIndex-1] = slice[tempCurIndex]
				slice[tempCurIndex] = oldLast
			}
		}

		return slice
	}

	return []int{}
}
