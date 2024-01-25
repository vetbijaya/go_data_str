package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	go func() {
		for _, value := range array {
			// ch <- &value => will print 4,4,4,4
			// Temporary variable v that has new memory address every iteration.
			v := value
			ch <- &v
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) // Output: 1,2,3,4
			wg.Done()
		}
	}()

	wg.Wait()
}
