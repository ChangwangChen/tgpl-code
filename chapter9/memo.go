package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"tgpl-code/chapter9/memo1"
	"tgpl-code/chapter9/memo2"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}


func main() {
	urls := []string{
		"http://www.baidu.com",
		"http://www.baidu.com",
		"http://www.qq.com",
		"http://www.qq.com",
		"http://www.163.com",
		"http://www.163.com",
		"http://www.163.com",
		"http://www.163.com",
		"http://www.163.com",
	}

	mem := memo1.New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := mem.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()

	fmt.Println("===============Memo2=======================")

	mem2 := memo2.New(httpGetBody)
	var n2 sync.WaitGroup
	for _, url1 := range urls {
		n2.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := mem2.Get(url1)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d\n",
				url, time.Since(start), len(value.([]byte)))
			n2.Done()
		}(url1)
	}
	n2.Wait()
	mem2.Close()

}
