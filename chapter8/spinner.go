package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\r Fibonacci(%d) = %d\n", n, fibN)
}

func fib(n int) int {
	if n < 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, x := range `-\|/` {
			fmt.Printf("\r%c", x)
			time.Sleep(delay)
		}
	}
}
