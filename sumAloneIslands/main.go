package main

import (
	"fmt"
)

// 找到值为1的元素，每次次数加1，将其周围为1的元素置为0，碰到0就返回
func searchIsland(plates [][]int) int {
	var num int
	for i := 0; i < len(plates); i++ {
		for j := 0; j < len(plates[i]); j++ {
			if plates[i][j] == 1 {
				spread(plates, i, j)
				num++
			}
		}
	}
	return num
}
func spread(plates [][]int, i, j int) {
	if i < 0 || i >= len(plates) || j < 0 || j >= len(plates[i]) || plates[i][j] == 0 {
		return
	}
	plates[i][j] = 0
	spread(plates, i+1, j)
	spread(plates, i-1, j)
	spread(plates, i, j+1)
	spread(plates, i, j-1)
}
func main() {
	plates := [][]int{
		{0, 0, 1, 1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 1},
		{1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
	}
	num := searchIsland(plates)
	fmt.Println(num)
}
