package bank_test

import (
	"testing"
	"TheGoProgrammingLanguage/ch9/bank1"
	"fmt"
)

func TestBank(t *testing.T){
	done := make(chan struct{})

	go func(){
		bank.Deposit(200)
		fmt.Println("=",bank.Balance())
		done <- struct{}{}
	}()

	go func(){
		bank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	//bank.Balance()

	if got, want := bank.Balance(), 300; got != want{
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
