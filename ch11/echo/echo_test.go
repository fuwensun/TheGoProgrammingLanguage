package main

import (
	"testing"
	"fmt"
	"bytes"
)

func TestEcho(t *testing.T){
	var tests = []struct{
		newline bool
		sep 	string
		args 	[]string
		want 	string
	}{
		{true,"",[]string{},"\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}

	for _, test := range tests{

		descr := fmt.Sprintf("echo(%v, %q, %q)",
			test.newline, test.sep, test.args)

		out = new(bytes.Buffer)
		if err := echo(test.newline, test.sep, test.args); err != nil{
			//t.Errorf<-------------
			t.Errorf("%s failed: %v",descr, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want{
			//t.Errorf<-------------
			//描述，结果，期望
			t.Errorf("--> %s = %q, want %q", descr, got, test.want)
		}


	}
	//fmt.Printf("-->%T-%T\n", out, new(string))
	//fmt.Printf("-->%T\n", out.(*bytes.Buffer))

}
