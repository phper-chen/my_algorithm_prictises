package main

import (
	"fmt"
)

func quickSort(arr []int, start, end int) {
	if start <= end {
		i, j := start, end
		mid := arr[(start+end)/2]
		for i <= j {
			for mid > arr[i] {
				i++
			}
			for mid < arr[j] {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func main() {
	arr := []int{23, 45, 3, 45, 78, 43, 2, 8, 2, 98, 234, 12, 7, 27, 9}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
