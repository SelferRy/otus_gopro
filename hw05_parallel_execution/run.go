package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	chTasks := buildTaskChannel(tasks)
	defer close(chTasks)
	goNum := minVal(n, len(tasks))
	err := readTaskChannel(chTasks, m, goNum)
	return err
}

func buildTaskChannel(tasks []Task) chan Task {
	ch := make(chan Task, len(tasks))
	go func() {
		for _, t := range tasks {
			ch <- t
		}
	}()
	return ch
}

func minVal(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func readTaskChannel(chTasks <-chan Task, errLim int, goNum int) error {
	var countErr int32
	done := make(chan struct{}) // filled if max err
	defer close(done)           // global (in main-scope) var
	var wg sync.WaitGroup
	for range goNum {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range chTasks {
				err := task()
				if err != nil {
					atomic.AddInt32(&countErr, 1)
				}
				if countErr >= int32(errLim) || len(chTasks) == 0 {
					return
				}
			}
		}()
	}
	wg.Wait()
	if countErr > 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
