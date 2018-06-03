package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//!+1
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-1

func main() {

	//初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//读到任意输入停止遍历
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		//我们可以将这个机制扩展一下，来作为我们的广播机制：
		// 不要向channel发送值，而是用关闭一个channel来进行广播。
		close(done)
	}()

	//遍历文件目录
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1) //<---------------
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait() //<---------------
		close(fileSizes)
	}()

	//打印结果
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			//排干file Size里的数据
			for range fileSizes {
				// Do nothing.
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	if cancelled() {
		return
	}

	defer n.Done() //<---------------
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1) //<---------------
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20) //超过20个go walkDir()后，这里会阻塞

func dirents(dir string) []os.FileInfo {

	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil
	}
	defer func() { <-sema }() // release token

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}
