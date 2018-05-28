package main

import (
	"os"
	"fmt"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte,1))
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown. Press return to abort.")
	select{
	case <- time.After(10 * time.Second):
		//do nothing
	case <- abort:
		fmt.Println("Launch aborted!")
		return

	}
	launch()
}

func launch(){
	fmt.Println("Lift off!")
}
