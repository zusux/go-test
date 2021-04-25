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
	limitCh := make(chan struct{},concurrencyProcesses)

	for i := 0; i < jobCount; i++ {
		limitCh <- struct{}{}
		go func(val int) {

			defer func() {
				wg.Done()
				<-limitCh
			}()
			waitTime := rand.Int31n(1000)
			fmt.Println("job:",val,"wait time:",waitTime,"millsecend")
			//time.Sleep(time.Duration(waitTime)*time.Second)
			time.Sleep(time.Second)
			found <- val
		}(i)
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