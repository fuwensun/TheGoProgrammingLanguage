package main

import (
	"io"
	"log"
	"net"
	"os"
)

//$ go build gopl.io/ch8/reverb1
//$ ./reverb1 &
//$ go build gopl.io/ch8/netcat2
//$ ./netcat2
//Hello?

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
