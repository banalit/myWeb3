package task2

import (
	"fmt"
	"sync"
)

func NumPrint() {
	wg := sync.WaitGroup{}
	printNum := func(odd bool) {
		wg.Add(1)
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if odd {
				if i%2 == 0 {
					fmt.Println("odd: ", i)
				}
			} else {
				if i%2 == 1 {
					fmt.Println("st: ", i)
				}
			}
		}
	}
	go printNum(true)
	go printNum(false)
	wg.Wait()
	fmt.Println("finish")

}
