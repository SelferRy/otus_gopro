package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	chErr := make(chan error, m)
	defer close(chErr)
	chTasks := buildTaskChannel(tasks)
	defer close(chTasks)
	goNum := minVal(n, len(tasks))
	err := readTaskChannel(chTasks, chErr, goNum)
	//var wg sync.WaitGroup
	//go buildTaskChannel(chTasks, tasks, &wg)
	return err
}

//func fillTaskChannel(ch chan<- Task, tasks []Task, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for _, t := range tasks {
//		ch <- t
//	}
//}

func buildTaskChannel(tasks []Task) chan Task { // ch chan<- Task, wg *sync.WaitGroup
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

func readTaskChannel(chTasks <-chan Task, chErr chan<- error, goNum int) error {
	done := make(chan struct{}) // filled if max err
	defer close(done)           // global (in main-scope) var
	var wg sync.WaitGroup
	for range goNum {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-done:
				return
			default:
				task := <-chTasks
				err := task()
				if err != nil {
					chErr <- err
				}
				if len(chErr) == cap(chErr) {
					done <- struct{}{}
				}
			}
		}()
	}
	wg.Wait()
	if len(chErr) > 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}

//func fillChannel(ch chan<- int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for t := range 4 {
//		ch <- t
//	}
//}

func fillChannel() chan int {
	n := 4
	ch := make(chan int, n)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for t := range n {
			ch <- t
		}
	}()

	go func() { // the func is infra part (only for waiting -> control runtime)
		wg.Wait()
		defer close(ch) // global (in main-scope) var
	}()

	return ch
}

func readChannel(ch <-chan int) {
	for result := range ch {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
