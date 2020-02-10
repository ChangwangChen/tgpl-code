package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8001")
	fmt.Println("dial localhost at 8001")
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	//
	//go mustCopy(os.Stdout, conn)
	//mustCopy(conn, os.Stdout)

	//下面是
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdout)
	conn.Close()
	<- done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
