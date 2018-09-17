package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

var contents = []string{
	"これは、私が小さいときに、村の茂平というおじいさんから聞いたお話です。",
	"昔は、私達の村の近くの、中山というところに小さなお城があって、",
	"中山様というお殿様が、おられたそうです。",
	"その中山から、少し離れた山の中に、「「ごんぎつね」という狐がいました。",
	"ごんは、一人ぼっちの子狐で、しだのいっぱい茂った森のなかに穴をほって住んでいました。",
	"そして、夜でも昼でも、当たりの村へ出てきて、イラズラばかりしました。",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		fmt.Fprintf(conn, stirngs.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contnets {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
			fmt.Fprintf(conn, "0\r\n\r\n")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:18888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
