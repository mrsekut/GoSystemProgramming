package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "ascii.jp:80")
		if err != nil {
			panic(err)
		}
		conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
		res, err := http.ReadResponse(bufio.NewReader(conn), nil)
		fmt.Println(res.Header)

		time.Sleep(3 * time.Second)
	}
}
