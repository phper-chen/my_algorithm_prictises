package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, aim int) int {
	low := 0
	height := len(arr) - 1
	for low <= height {
		var mid int = low + (height-low)/2
		var midV int = arr[mid]
		if midV == aim {
			return midV
		} else if midV > aim {
			height = mid - 1
		} else if midV < aim {
			low = mid + 1
		}
	}
	return -1
}
func main() {
	arr := []int{2, 4, 5, 7, 3, 9, 2, 3, 56, 23}
	sort.Ints(arr)
	re := binarySearch(arr, 4)
	fmt.Println(re)
}
