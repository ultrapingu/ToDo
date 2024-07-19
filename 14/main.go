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

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < max; i += 2 {
			num = i
			fmt.Printf("[%d] %d\n", i, num)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < max; i += 2 {
			num = i
			fmt.Printf("[%d] %d\n", i, num)
		}
	}()

	wg.Wait()
}
