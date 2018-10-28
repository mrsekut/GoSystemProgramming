package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("task 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("task 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("end")
}
