package main

import "fmt"

//单向通道示例
func counter(out chan<- int) {
	for i := 0; i <= 100; i++ {
		out <- i
	}
	//发送方才能关闭一个仅能接收的 chan
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}
