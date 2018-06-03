package main

import (
	"io"
	"log"
	"net"
	"time"
)

//nc localhost 8000
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		//_, err := io.WriteString(c, time.Now().Format("03:04:05\n"))
		//_, err := io.WriteString(c, time.Now().Format("03:04:05PM\n"))
		//_, err := io.WriteString(c, time.Now().Format("2006 03:04:05PM\n"))
		_, err := io.WriteString(c, time.Now().Format("Mon Jan 2 03:04:05PM 2006 UTC-0700\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
