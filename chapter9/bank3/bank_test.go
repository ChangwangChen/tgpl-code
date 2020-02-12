package bank_test

import (
	"fmt"
	"sync"
	"testing"
	bank "tgpl-code/chapter9/bank3"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}

	n.Wait()

	if got, want := bank.Balance(), (1000 + 1) * 1000 / 2; got != want {
		t.Errorf("Balance = %d, want = %d", got, want)
	}

	fmt.Println("Balance:", bank.Balance())

	if !bank.Withdraw(1000) {
		t.Errorf("Balance = %d, want > %d", bank.Balance(), 1000)
	}

	fmt.Println("Balance:", bank.Balance())
}
