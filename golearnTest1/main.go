package main

import "fmt"

func main() {
	var numbers []int = []int{1, 2, 2, 3, 4, 3, 4}
	single := singleNumber(numbers)
	fmt.Println("single is: ", single)
}

// 单一个数字
func singleNumber(numbers []int) int {
	single := 0
	for _, num := range numbers {
		single ^= num
	}
	return single
}
