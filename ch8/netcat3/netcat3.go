package main

import (
	"log"
	"io"
	"os"
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	fmt.Println("----")
	mustCopy(conn,os.Stdin)
	fmt.Println("++++")
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader){

	//！！！要点！！！
	// 阻塞式的从src输入流读数据到dst输出流
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
