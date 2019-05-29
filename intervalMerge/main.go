package main

import (
	"fmt"
	"sort"
)

//合并重叠区间 给定一组 区间，合并所有重叠的 区间。

//给定：[1,3],[2,6],[8,10],[15,18] 返回：[1,6],[8,10],[15,18]

type Interval struct {
	start int
	end   int
}

func merge(intervals []Interval) []Interval {
	lens := len(intervals)
	if lens == 0 || lens == 1 {
		return intervals
	}
	// 给切片排序 按start从小到大的顺序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	fmt.Println(intervals)
	// 新的结果切片
	newIntervals := make([]Interval, 0)
	swap := Interval{}
	for k, v := range intervals {
		if k == 0 {
			swap = v
			continue
		}
		fmt.Println(len(newIntervals))
		fmt.Println(cap(newIntervals))
		// 必须将第一个作为参考点来开启逻辑，如果第一次相交或者重合就改变结束区间
		if v.start < swap.end {
			if v.end < swap.end {
				continue
			}
			swap.end = v.end
			newIntervals = append(newIntervals, swap)
		} else {
			// 不相交则加入到新的结果切片里
			newIntervals = append(newIntervals, v)
			// 将该次的点作为参考点
			swap = v
		}
	}
	return newIntervals
}

func main() {
	a := []Interval{
		{14, 30}, {1, 4}, {3, 8}, {9, 11}, {12, 45},
	}
	b := merge(a)
	fmt.Println(b)
}
