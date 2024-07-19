package main

import (
	"fmt"
)

func main() {
	todoItems := []string{"item 1", "item 2", "item 3"}
	todostatuses := []string{"status 1", "status 2", "status 3"}

	itemsCh := make(chan string)
	resultsCh := make(chan string)

	go func() {
		for _, item := range todoItems {
			itemsCh <- item
		}
		close(itemsCh)
	}()

	go func() {
		for _, status := range todostatuses {
			item := <-itemsCh
			resultsCh <- fmt.Sprintf("[%s, %s]", item, status)
		}
		close(resultsCh)
	}()

	for result := range resultsCh {
		fmt.Println(result)
	}
}
