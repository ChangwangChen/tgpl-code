package bank

var (
	sema    = make(chan struct{}, 1)
	balance int //包私有， 外界不能直接访问
)

func Deposit(amount int) {
	sema <- struct{}{} //获取
	balance += amount
	<-sema //释放
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
