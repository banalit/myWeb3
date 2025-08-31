package main

import (
	"fmt"
	"sort"
)

// 合并区间
func main() {
	// 测试案例
	testCases := [][]Interval{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{1, 10}, {2, 3}, {5, 8}, {11, 15}},
		{{}}, // 空区间
	}

	// 执行测试并输出结果
	for i, tc := range testCases {
		fmt.Printf("测试案例 %d:\n", i+1)
		fmt.Printf("  原始区间: %v\n", tc)
		fmt.Printf("  合并后:   %v\n\n", mergeIntervals(tc))
	}
}

// Interval 表示一个区间
type Interval struct {
	Start int
	End   int
}

// 合并区间的函数
func mergeIntervals(intervals []Interval) []Interval {
	// 处理空输入
	if len(intervals) == 0 {
		return nil
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		}
		if (intervals[i].Start == intervals[j].Start) && (intervals[i].End < intervals[j].End) {
			return true
		}
		return false

	})

	// 初始化结果切片，放入第一个区间
	merged := []Interval{intervals[0]}

	// 遍历剩余区间并合并
	for i := 1; i < len(intervals); i++ {
		// 获取结果切片的最后一个区间
		last := &merged[len(merged)-1]

		// 如果当前区间与最后一个区间重叠或相邻，则合并它们
		if intervals[i].Start <= last.End {
			// 合并后的区间结束位置取较大值
			if intervals[i].End > last.End {
				last.End = intervals[i].End
			}
		} else {
			// 如果不重叠，则直接添加到结果切片
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
