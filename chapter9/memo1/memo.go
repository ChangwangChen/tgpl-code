package memo1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

//缓存 func 调用的结果
type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

//Memo 使用的 func 原型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{} //res 准备好之后直接关闭
}

func New(f Func) *Memo {
	return &Memo{
		f: f,
		cache: make(map[string]*entry),
	}
}

func (mem *Memo) Get(key string) (interface{}, error) {
	mem.mu.Lock()
	e := mem.cache[key]
	if e == nil {
		fmt.Println("URL: ", key)
		e = &entry{ready: make(chan struct{})}
		mem.cache[key] = e
		mem.mu.Unlock() //写入之后就直接解锁， 下次就直接不会走这个分支

		e.res.value, e.res.err = mem.f(key)
		close(e.ready) //关闭 ready chan
	} else {
		mem.mu.Unlock()
		//当第一个 goroutine 在执行的时候， 不会关闭 ready 的渠道， 这里会一直阻塞
		//当关闭了 ready 通道的时候， 下面的语句会直接返回
		<-e.ready //等待数据完毕
	}

	return e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

