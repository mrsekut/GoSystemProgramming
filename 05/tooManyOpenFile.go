package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	for i := 1; i < 1000; i++ {
		path := "many_files/test_" + strconv.Itoa(i) + ".txt"
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		file.Write([]byte("os.File example\n"))
	}

	time.Sleep(30 * time.Second)
}
