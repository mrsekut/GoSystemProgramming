// sync.Mutexを使うと、「メモリを読み込んで書き換えるコード」にはいるgoroutineを一つに制限できる
package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	id++
	result := id
	mutex.Unlock()
	return result
}

func main() {

	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}