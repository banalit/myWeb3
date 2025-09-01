package main

import (
	"fmt"

	"github.com/luke/web3Learn/task2"
)

func main() {
	fmt.Println("go rountine task start")
	task2.NumPrint()
	task2.SchedulerTest()

	fmt.Println("\npointer task start")
	num := []int{200, 100, 300}
	task2.PtrAdd10(&num[0])
	task2.SliceMulti2(&num)

	fmt.Println("\ninterface object test")
	task2.ShapeTest()
	fmt.Println("\nemployee test")
	task2.TestEmployee()

	fmt.Println("\nchanneltest")
	task2.ChannelTest()
	task2.ChannelTest2()

	fmt.Println("\nlockTest")
	task2.AtomicTest()
	task2.LockTest()

}
