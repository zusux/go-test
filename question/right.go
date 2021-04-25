package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	const concurrencyProcesses = 10
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(jobCount)
	found := make(chan int)
	//limitCh := make(chan struct{},concurrencyProcesses)
	queue := make(chan int)

	go func(queue chan<- int) {
		for i := 0; i < jobCount; i++ {
			queue <- i
		}
		close(queue)
	}(queue)


	for i := 0; i < concurrencyProcesses; i++ {
		go func(queue <-chan int) {
			for val := range queue{
				defer  wg.Done()
				waitTime := rand.Int31n(1000)
				fmt.Println("job:",val,"wait time:",waitTime,"millsecend")
				time.Sleep(time.Duration(waitTime)*time.Millisecond)
				time.Sleep(time.Second)
				found <- val
			}
		}(queue)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	var result []int
	for p := range found{
		fmt.Println("finished job:",p)
		result = append(result,p)
	}
	fmt.Println("result:",result)
}