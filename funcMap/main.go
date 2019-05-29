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

var funcMap = map[string]func(n int) int{
	"a": test1,
	"b": test2,
	"c": test3,
}

func main() {
	arr := []string{"a", "b", "c"}
	for i, v := range arr {
		funcMap[v](i)
	}
}
