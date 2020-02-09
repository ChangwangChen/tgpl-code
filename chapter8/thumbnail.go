package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func makeThumbnails(counts int) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for i:=0; i< counts;i++ {
		wg.Add(1)
		//worker
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			sizes <- rand.Int63n(1000)
		}()
	}

	//closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	//sum
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	fmt.Println("sum: ", makeThumbnails(100))
}
