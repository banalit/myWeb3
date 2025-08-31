package main

import "fmt"

// 删除有序数字数组中重复的数字
func main() {
	test := []int{1, 1, 2}
	removeDuplicate(test)
	test = []int{1, 1, 2, 2}
	removeDuplicate(test)
	test = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicate(test)

}

func removeDuplicate(test []int) []int {
	if len(test) < 2 {
		fmt.Println(test, test)
		return test
	}
	var after []int
	for i := 0; i < len(test)-1; i++ {
		if test[i] == test[i+1] {
			continue
		}
		after = append(after, test[i])
	}
	after = append(after, test[len(test)-1])
	fmt.Println(test, after)
	return after
}
