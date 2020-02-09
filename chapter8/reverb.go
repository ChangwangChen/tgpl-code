package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn1(conn) //每个链接都会开启一个新的 goroutine, 但是处理每个链接的请求是顺序进行的
	}
}

func handleConn1(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, s string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}
