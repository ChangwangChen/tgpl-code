package memo2

type request struct {
	key string
	response chan<- result
}

//缓存 func 调用的结果
type Memo struct {
	requests chan request
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
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (mem *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	mem.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (mem *Memo) Close() {
	close(mem.requests)
}

func (mem *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range mem.requests {
		e := cache[req.key]
		if e == nil {
			//
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

