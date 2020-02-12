package bank

import "sync"

var (
	mutex   sync.Mutex
	balance int
)

func Deposit(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	deposit(amount)
}

func Balance() int {
	mutex.Lock()
	defer mutex.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mutex.Lock()
	defer mutex.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}
