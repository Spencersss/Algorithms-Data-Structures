package main

import (
	"testing"
)

// Test vars
var testTree []int

// Setup values for testing
func init() {
	testTree = []int{-3, -2, -2, 0, 0, 1, 3, 4, 5, 6, 15, 16}
}

// Tests
func TestSuccessfulBinarySearchIterative(t *testing.T) {
	result := BinarySearchIterative(testTree, 5)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d but got %d from BinarySearchIterative.", expected, result)
	}
}

func TestNotFoundValueBinarySearchIterative(t *testing.T) {
	result := BinarySearchIterative(testTree, 12412312)
	expected := -1

	if result != expected {
		t.Errorf("Expected %d but got %d from BinarySearchIterative.", expected, result)
	}
}

func TestSuccessfulBinarySearchRecursive(t *testing.T) {
	result := BinarySearchRecursive(testTree, -2)
	expected := 2

	if result != expected {
		t.Errorf("Expected %d but got %d from BinarySearchIterative.", expected, result)
	}
}

func TestNotFoundValueBinarySearchRecursive(t *testing.T) {
	result := BinarySearchRecursive(testTree, 6123123)
	expected := -1

	if result != expected {
		t.Errorf("Expected %d but got %d from BinarySearchIterative.", expected, result)
	}
}
