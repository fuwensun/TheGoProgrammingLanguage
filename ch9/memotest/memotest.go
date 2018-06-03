package memotest

import (
	"net/http"
	"io/ioutil"
	"testing"
	"time"
	"log"
	"fmt"
	"sync"
)

func httpGetBody(url string)(interface{}, error){
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string{
	ch := make(chan string)
	go func(){
		for _, url := range []string{
			"http://www.github.com",
			"http://www.youku.com",
			"http://www.baidu.com",
			"http://www.sina.com",
		}{
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface{
	Get(key string)(interface{}, error)
}

func Sequential(t *testing.T, m M){
	for url := range incomingURLs(){
		start := time.Now()
		value, err := m.Get(url)
		if err != nil{
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M){
	var n sync.WaitGroup
	for url := range incomingURLs(){
		n.Add(1)
		go func(url string){
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil{
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}
