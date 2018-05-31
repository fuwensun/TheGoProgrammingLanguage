package bank

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int){
	deposits <- amount
}

func Balance()int{
	return <- balances
}

func teller(){
	var balance int
	for{
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			fmt.Printf("balances <- %d\n",balance)
		}
	}
}

func init(){
	go teller()
}