package main

import (
	"fmt"
	"sync"
)

var num int

func main() {
	const max = 100
	num = 0

	var wg sync.WaitGroup
	results := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < max; i += 2 {
			results <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < max; i += 2 {
			results <- i
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		num = result
		fmt.Println(num)
	}
}
