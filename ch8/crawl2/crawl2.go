package main

import (
	"fmt"
	"TheGoProgrammingLanguage/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string)[]string{
	fmt.Println(url)

	tokens <- struct{}{}			//<--------
	list,err := links.Extract(url)
	<- tokens						//<--------

	if err != nil{
		log.Print(err)
	}
	return list
}


//./crawl2.exe https://www.github.com

func main() {
	worklist := make(chan []string)
	var n int //number of pending sends to worklist

	n++										//<--------
	go func(){worklist <- os.Args[1:]}()

	seen := make(map[string]bool)

	for; n > 0; n--{
		list := <- worklist					//<--------
		for _, link := range list{
			if !seen[link]{
				seen[link] = true

				n++							//<--------
				go func(link string){
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
