package binarysearch

import (
	"slices"
)

// SearchInts useing binary search algorithm
func SearchInts(list []int, key int) int {
	slices.Sort(list)
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2 // Find the midpoint.

		// Check if the midpoint is the key.
		if list[mid] == key {
			return mid // Found the key.
		}

		// If the key is smaller, ignore the right half.
		if list[mid] > key {
			high = mid - 1
		} else { // If the key is larger, ignore the left half.
			low = mid + 1
		}
	}
	return -1
}
