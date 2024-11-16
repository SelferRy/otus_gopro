package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n int, m int) error {
	if n <= 0 {
		return errors.New("n must be > 0")
	}

	if m <= 0 {
		m = len(tasks) + 1
	}

	tasksChan := make(chan Task)
	var errorCount int32

	wg := sync.WaitGroup{}
	goNum := minVal(n, len(tasks))
	listenChannel(tasksChan, &errorCount, goNum, &wg)
	fillChannel(tasksChan, &errorCount, tasks, m)
	wg.Wait()

	if errorCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func fillChannel(tasksChan chan<- Task, errorCount *int32, tasks []Task, m int) {
	defer close(tasksChan)
	for _, task := range tasks {
		if atomic.LoadInt32(errorCount) >= int32(m) {
			break
		}
		tasksChan <- task
	}
}

func listenChannel(tasksChan <-chan Task, errorCount *int32, goNum int, wg *sync.WaitGroup) {
	for range goNum {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChan {
				if task() != nil {
					atomic.AddInt32(errorCount, 1)
				}
			}
		}()
	}
}

func minVal(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
