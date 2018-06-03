package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("\n\n---array------------")
	a := [...]string{0: "a", 1: "b"}
	a = [...]string{"a", "b"}
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("数组类型！！！--> %T\n", a)
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(&a[1])

	fmt.Println("\n\n---slice-----------")
	s := a[:len(a)]
	fmt.Println(reflect.TypeOf(s))
	fmt.Printf("切片类型！！！--> %T\n", s)
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(&s[1])

	fmt.Println("----s1-----")
	s = make([]string, 2, 4)
	fmt.Printf("%T\n", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("----s2-----")
	s = []string{"a", "b"}
	fmt.Printf("%T\n", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("----s3-----")
	s = make([]string, 2)
	fmt.Printf("%T\n", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("----s4-----")
	s = make([]string, 4)[:2]
	fmt.Printf("%T\n", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("\n\n---map------------")
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)
	fmt.Println(m["a"])
	//fmt.Println(&m["a"]) //map不能取地址 <---------

}
