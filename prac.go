package main

import (
	"os"
	"time"
)

func renameTest() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	file2, err := os.Create("sample2.txt")
	if err != nil {
		panic(err)
	}
	file3, err := os.Create("sample3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer file2.Close()
	defer file3.Close()

	time.Sleep(60 * time.Second)

}

func main() {
	renameTest()
}
