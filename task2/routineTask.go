package task2

import (
	"fmt"
	"sync"
	"time"
)

func NumPrint() {
	wg := sync.WaitGroup{}
	printNum := func(odd bool) {
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
	// printNum(true)
	wg.Add(2)
	go printNum(true)
	go printNum(false)
	wg.Wait()

}

type Task struct {
	Name string
	Cost int64
	Func func()
}
type TaskResult struct {
	TaskName string
	Duration time.Duration
}

type TaskScheduler struct {
	tasks       []Task
	taskResults []TaskResult
	mutex       sync.Mutex
	wg          sync.WaitGroup
}

func (s *TaskScheduler) AddTask(t Task) {
	s.tasks = append(s.tasks, t)
}

func (s *TaskScheduler) Run() {
	s.wg.Add(len(s.tasks))

	for _, t := range s.tasks {
		go func(task Task) {
			defer s.wg.Done()
			st := time.Now()
			task.Func()

			since := time.Since(st)

			s.mutex.Lock()
			tr := TaskResult{
				TaskName: task.Name,
				Duration: since,
			}
			s.taskResults = append(s.taskResults, tr)
			s.mutex.Unlock()
		}(t)
	}

	s.wg.Wait()
}

func (s *TaskScheduler) PrintResults() {
	fmt.Println("任务执行结果：")
	fmt.Println("---------------------")
	for _, result := range s.taskResults {
		fmt.Printf("任务 %-10s 执行时间: %v\n", result.TaskName, result.Duration)
	}
}

func SchedulerTest() {
	// 创建任务调度器
	scheduler := TaskScheduler{}

	// 添加示例任务
	scheduler.AddTask(Task{
		Name: "任务1",
		Func: func() {
			time.Sleep(100 * time.Millisecond) // 模拟任务执行
			fmt.Println("任务1执行完毕")
		},
	})

	scheduler.AddTask(Task{
		Name: "任务2",
		Func: func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("任务2执行完毕")
		},
	})

	scheduler.AddTask(Task{
		Name: "任务3",
		Func: func() {
			time.Sleep(150 * time.Millisecond)
			fmt.Println("任务3执行完毕")
		},
	})

	// 执行所有任务
	startTime := time.Now()
	scheduler.Run()
	totalDuration := time.Since(startTime)

	// 打印结果
	scheduler.PrintResults()
	fmt.Printf("---------------------\n")
	fmt.Printf("总执行时间: %v\n", totalDuration)
}
