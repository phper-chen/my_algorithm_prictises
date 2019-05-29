package main

import (
	//	"bytes"
	"fmt"
)

func sumLineSegsLen(arr [][]int) int {
	points := map[int]int{}

	for i := 0; i < len(arr); i++ {
		start, end := arr[i][0], arr[i][1]
		if start > end {
			start, end = end, start
		}
		fmt.Printf("start is %d, end is %d\n", start, end)
		for j := start; j <= end; j++ {
			points[j] = 1
		}
	}
	count := 0
	fmt.Printf("%v", points)

	for range points {
		count++
	}
	fmt.Println(count)
	return count
}
func main() {
	arr := [][]int{
		{1, 3}, {2, 4}, {7, 15}, {19, 17},
	}
	num := sumLineSegsLen(arr)
	fmt.Println(num)
}
