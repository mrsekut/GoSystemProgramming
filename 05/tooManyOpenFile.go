package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	for {
		data, err := ioutil.ReadFile("test.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
