package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func main() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}

	fmt.Println("-->" + cwd)
}
