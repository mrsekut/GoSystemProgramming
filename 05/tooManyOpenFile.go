package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	for i := 1; i < 1000; i++ {
		path := "many_files/test_" + strconv.Itoa(i) + ".txt"
		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}

	time.Sleep(30 * time.Second)
}
