package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//无缓冲通道
	naturals := make(chan int)
	squares := make(chan int)

	fmt.Println("chan int: ", unsafe.Sizeof(naturals), " bytes")

	//counter
	go func() {
		// 100 个数依次写入 naturals
		for i := 0; i <= 100; i++ {
			naturals <- i
		}

		//关闭 chan
		close(naturals)
	}()

	//squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}

		close(squares)
	}()

	//printer1
	for {
		//这里使用 ok 来判断 chan squares 是否已经关闭并且读取完毕
		//因为从一个已经关闭的 chan 读取数据的时候， 始终会返回 chan Type 的默认值
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}

	//printer2
	for x := range squares {
		fmt.Println(x)
	}
}
