package main

import (
	"fmt"
)

func test1(n int) int {
	fmt.Println("test1")
	return n + 1
}
func test2(n int) int {
	fmt.Println("test2")
	return n + 2
}
func test3(n int) int {
	fmt.Println("test3")
	return n + 3
}

var fucArr = []func(n int) int{test1, test2, test3}

func main() {
	for i, v := range "abc" {
		fmt.Println(v)
		fucArr[i](i)
	}
}
