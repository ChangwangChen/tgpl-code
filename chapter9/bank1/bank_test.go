package bank_test

import (
	"fmt"
	"testing"

	"tgpl-code/chapter9/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposits(100)
		fmt.Println("Balance:", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposits(200)
		fmt.Println("Balance:", bank.Balance())
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
