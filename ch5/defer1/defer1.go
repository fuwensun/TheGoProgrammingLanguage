package main

import "fmt"

func main() {
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//正常输出 <---------
//f(3)
//f(2)
//f(1)

//函数defer在panic运行前，被调用并输出 <---------
//defer 1
//defer 2
//defer 3

//panic运行并输出 <---------
//panic: runtime error: integer divide by zero
//
//goroutine 1 [running]:
//main.f(0x0)
//C:/Users/sfw-sp/go/src/TheGoProgrammingLanguage/ch5/my1/main.go:9 +0x1a4
//main.f(0x1)
//C:/Users/sfw-sp/go/src/TheGoProgrammingLanguage/ch5/my1/main.go:11 +0x173
//main.f(0x2)
//C:/Users/sfw-sp/go/src/TheGoProgrammingLanguage/ch5/my1/main.go:11 +0x173
//main.f(0x3)
//C:/Users/sfw-sp/go/src/TheGoProgrammingLanguage/ch5/my1/main.go:11 +0x173
//main.main()
//C:/Users/sfw-sp/go/src/TheGoProgrammingLanguage/ch5/my1/main.go:6 +0x31
//
//Process finished with exit code 2
