package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math"
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, t := range tasks {
			ch <- t
		}
	}()
	wg.Wait()
	return ch
}

func minVal(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func readTaskChannel(chTasks <-chan Task, errLim int, goNum int) error {
	var (
		countErr int32
		wg       sync.WaitGroup
	)
	if errLim >= math.MaxInt32 {
		return fmt.Errorf("the maximum number of allowed errors has been exceeded")
	}
	for range goNum {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range chTasks {
				err := task()
				if err != nil {
					atomic.AddInt32(&countErr, 1)
				}
				//nolint:gosec
				if atomic.LoadInt32(&countErr) >= int32(errLim) || len(chTasks) == 0 {
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
