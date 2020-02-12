package bank_test

import (
	"sync"
	"testing"
	bank "tgpl-code/chapter9/bank2"
)

func TestBank(t *testing.T) {
	var wg sync.WaitGroup

	var amount int
	for i := 0; i < 1000; i++ {
		amount += i
		wg.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if got := bank.Balance(); got != amount {
		t.Errorf("Balance = %d, want = %d", got, amount)
	}
}
