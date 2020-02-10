package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string //单向通道

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				if clients[cli] { //这里只向正常存在的客户端发送消息
					cli <- msg
				}
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			clients[cli] = false
			close(cli)
		}
	}
}

func handleConn2(c net.Conn) {
	ch := make(chan string)
	go clientWriter(c, ch)

	who := c.RemoteAddr().String()
	fmt.Println("who: " + who)
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- who + " says: " + input.Text()
	}

	//input EOF
	leaving <- ch
	messages <- who + " has left "
	c.Close()
}

func clientWriter(c net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn2(conn)
	}
}
