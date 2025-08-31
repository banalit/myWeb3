package main

import "fmt"

// 两数之和, 返回index
func main() {
	var test []int = []int{1, 2, 3, 4, 5}
	var target = 8
	fmt.Printf("%v target %v, idx: %v", test, target, getArrayIndex(test, target))
}

func getArrayIndex(test []int, target int) []int {
	for i := 0; i < len(test)-1; i++ {
		for j := i + 1; j < len(test); j++ {
			if test[i]+test[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
