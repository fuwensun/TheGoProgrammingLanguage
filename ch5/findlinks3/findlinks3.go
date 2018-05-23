package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

func breadthFirst(f func(item string)[]string, worklist []string){
	seen := make(map[string]bool)
	for len(worklist) > 0{  // <-------1,判断worklist有内容
		items := worklist
		worklist = nil		// <-------2,清零worklist
		for _, item := range items{
			if !seen[item]{
				seen[item] = true
				worklist = append(worklist, f(item)...)	// <-------3,从新填充worklist
			}
		}
	}
}

func crawl(url string)[]string{
	fmt.Println("-->" + url)
	list, err := links.Extract(url)
	if err != nil{
		log.Print(err)
	}
	return list
}
func main() {
	breadthFirst(crawl, os.Args[1:])
}
