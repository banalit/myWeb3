package task2

import (
	"fmt"
	"sync"
)

/*
*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*
*/
func ChannelTest() {
	fmt.Println("ChannelTest:")
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		produce(ch, 10)

	}()
	go func() {
		defer wg.Done()
		consume(ch)
	}()
	wg.Wait()

}

func produce(ch chan<- int, size int) {
	for i := 1; i <= size; i++ {
		// fmt.Println("before add ", i)
		ch <- i
		fmt.Println("after add ", i)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for i := range ch {
		fmt.Println("get int :", i)
	}
}

/*
*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*
*/
func ChannelTest2() {
	fmt.Println("ChannelTest2:")
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 10)
	go func() {
		defer wg.Done()
		produce(ch, 100)

	}()
	go func() {
		defer wg.Done()
		consume(ch)
	}()
	wg.Wait()

}
