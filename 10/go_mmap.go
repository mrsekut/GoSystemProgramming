package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	mmap "github.com/edsrzf/mmap-go"
)

func main() {
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 指定したファイルの内容をメモリ上に展開
	m, err := mmap.Map(f, mmap.COPY, 0) // mmap.COPYにしてみる
	if err != nil {
		panic(err)
	}
	defer m.Unmap()

	m[9] = 'X'
	// m.Flush() // 書きかけの内容をファイルに保存

	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original:  %s\n", testData)
	fmt.Printf("mmap: 	   %s\n", m)
	fmt.Printf("file: 	   %s\n", fileData)
}
