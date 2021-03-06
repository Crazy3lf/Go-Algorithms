package searches

import "errors"

var errNotFound = errors.New("target not found in array")

// BinarySearch Binary Search
func BinarySearch(array []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, errNotFound
	}
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return BinarySearch(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return BinarySearch(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

// IterBinarySearch Iter Binary Search
func IterBinarySearch(array []int, target int, lowIndex int, highIndex int) (int, error) {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, errNotFound
}
