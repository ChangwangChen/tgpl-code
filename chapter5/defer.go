package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)                   //panic if x == 0
	defer func() { fmt.Printf("defer %d\n", x) }() //匿名函数
	f(x - 1)                                       //递归
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	defer printStack()
	f(4)
}
