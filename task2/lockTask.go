package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*
*/
type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func LockTest() {
	var wg = sync.WaitGroup{}
	var counter Counter
	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("lock counter:", counter.count)

}

/*
*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*
*/
func AtomicTest() {
	var wg = sync.WaitGroup{}
	var counter int64
	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("atomic counter:", counter)

}
