package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
goroutine 1 [running]:
main.printStack()//后 找到 defer 函数, 打印全部栈(包括自己和panic)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:16 +0x5b
panic(0x4a20c0, 0x554270)//先 panic
        /usr/local/go/src/runtime/panic.go:522 +0x1b5
main.f(0x0)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:21 +0x174
main.f(0x1)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:23 +0x14f
main.f(0x2)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:23 +0x14f
main.f(0x3)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:23 +0x14f
main.main()
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:11 +0x46
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x0)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:21 +0x174
main.f(0x1)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:23 +0x14f
main.f(0x2)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:23 +0x14f
main.f(0x3)
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:23 +0x14f
main.main()
        /mnt/i/github.com/gkt_cc_go/src/example/defer2/defer2.go:11 +0x46
exit status 2
*/
