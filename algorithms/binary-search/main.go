package main

import "fmt"

func main() {

	// Initialize a slice representing a binary tree
	tree := []int{-3, -2, -2, 0, 0, 1, 3, 4, 5, 6, 15, 16}

	// Perform iterative binary search
	result := BinarySearchIterative(tree, 5)

	// Print result
	fmt.Println(result)

	// Perform recursive binary search
	result = BinarySearchRecursive(tree, 5)

	// Print result
	fmt.Println(result)
}

// BinarySearchIterative is a iterative implementation of the binary search algorithm. Returns the index of the value
// if found and returns -1 if the value is not found.
func BinarySearchIterative(tree []int, value int) int {
	low := 0
	high := len(tree) - 1

	for low <= high {

		mid := low + (high-low)/2
		if tree[mid] == value {
			return mid
		}

		if value < tree[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchRecursive is a recursive implementation of the binary search algorithm. Returns the index of the value
// if found and returns -1 if the value is not found.
func BinarySearchRecursive(tree []int, value int) int {
	return binarySearchRecursiveHelper(tree, value, 0, len(tree)-1)
}

// function to assist with passing both the low and high bounds when traversing the binary search algorithm
// allows both the iterative and recursive function definitions to maintain identical parameters.
func binarySearchRecursiveHelper(tree []int, value int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if tree[mid] == value {
		return mid
	}

	if value < tree[mid] {
		return binarySearchRecursiveHelper(tree, value, low, mid)
	} else {
		return binarySearchRecursiveHelper(tree, value, mid+1, high)
	}
}
