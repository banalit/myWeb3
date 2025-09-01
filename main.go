package main

import (
	"fmt"

	"github.com/luke/web3Learn/task2"
)

func main() {
	fmt.Println("go rountine task start")
	task2.NumPrint()
	task2.SchedulerTest()

	fmt.Println("pointer task start")
	num := []int{200, 100, 300}
	task2.PtrAdd10(&num[0])
	task2.SliceMulti2(&num)
}
