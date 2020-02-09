package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	//接收到取消信息， 直接返回
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case <-done:
		return nil
	case sema <- struct{}{}:
	}

	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)
var done = make(chan struct{}) //取消信号的 通道

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."} //默认当前目录
	}

	//启用新 goroutine 监听 StdIn， 接受取消信息
	go func() {
		os.Stdin.Read(make([]byte, 1)) //这里比较简陋， 任意的一个输入都当作取消的信息
		close(done) //关闭 done chan 当作取消的信息
	}()

	//遍历
	fileSizes := make(chan int64)
	//go func() {
	//	for _, root := range roots {
	//		walkDir(root, fileSizes)
	//	}
	//	close(fileSizes)
	//}()

	//作用是在等待所有 goroutine 执行完成
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait() //等待所有的 walkDir 执行完成之后， 关闭 fileSizes chan
		close(fileSizes)
	}()

	//定期输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(50 * time.Millisecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case <-done:
			for range fileSizes {
				//收到取消信号之后， 读取完所有 fileSizes chan 中的信息， 防止阻塞 goroutine
			}
			return
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(n, s int64) {
	fmt.Printf("%d files  %.2f GB\n", n, float64(s)/1e9)
}
