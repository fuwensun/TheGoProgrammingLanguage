package main

import "fmt"

var pc [10]byte

//func init() {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//}

//var pc [10]byte = func() (pc [10]byte) {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//	return
//}()

func main() {
	for i, v := range pc {
		fmt.Printf("%d --> %v\n", i, v)
	}

	fmt.Println()
	for i := range pc {
		fmt.Printf("%d --> %v\n", i, i)
	}

}
