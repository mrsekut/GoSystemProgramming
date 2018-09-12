package main

import (
	"os"
	"time"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	time.Sleep(30 * time.Second)
}
