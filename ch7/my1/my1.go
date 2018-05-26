package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
)
func main(){

}
func init() {

	//一、	第一种，如果断言的类型T是一个具体类型，然后类型断言检查x的动态类型是否和T相同。
	//如果这个检查成功了，类型断言的结果是x的动态值，当然它的类型是T。
	//换句话说，具体类型的类型断言从它的操作对象中获得具体的值。
	//var w io.Writer
	//w = os.Stdout
	//f := w.(*os.File)      // success: f == os.Stdout
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f.Name())

	//c := w.(*bytes.Buffer)	// panic: interface holds *os.File, not *bytes.Buffer
}

func init(){

	//二、	第二种，如果相反断言的类型T是一个接口类型，然后类型断言检查是否x的动态类型满足T。
	//如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值，但是结果有类型T。
	//换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），
	//但是它保护了接口值内部的动态类型和值的部分。
	//
	//在下面的第一个类型断言后，w和rw都持有os.Stdout因此它们每个有一个动态类型*os.File，
	//但是变量w是一个io.Writer类型只对外公开出文件的Write方法，然而rw变量也只公开它的Read方法。
	//
	//var w io.Writer
	//w = os.Stdout
	//rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	//w = new(ByteCounter)

	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write

	//w = new(ByteCounter)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method


	//三、	我们几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言，
	//因为它表现的就像赋值操作一样，除了对于nil接口值的情况。

	w = rw             // io.ReadWriter is assignable to io.Writer
	w = rw.(io.Writer) // fails only if rw == nil
}

func init(){
	//如果类型断言出现在一个预期有两个结果的赋值操作中，例如如下的定义，
	//这个操作不会在失败的时候发生panic但是代替地返回一个额外的第二个结果，
	//这个结果是一个标识成功的布尔值：

	var w io.Writer = os.Stdout
	f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
	fmt.Printf("\n%v-%v\n",f.Name(),ok)
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	fmt.Printf("%v-%v\n",b,ok)

	if f, ok := w.(*os.File); ok {
		// ...use f...
		fmt.Printf("%v-%v\n",f.Name(),ok)
	}

	if w, ok := w.(*os.File); ok {
		// ...use w...
		fmt.Printf("%v-%v\n",w.Name(),ok)
	}
}


type MyReader interface {
	Read(p []byte) (n int, err error)
}
type MyWriter interface {
	Write(p []byte) (n int, err error)
}
type MyReadWriter interface {
	MyReader
	MyWriter
}

type MyNewReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

func init(){

	var rw MyReadWriter = os.Stdout
	r := rw.(MyReader)
	w := rw.(MyWriter)
	fmt.Printf("\n%T - %T - %T\n",rw,r,w)
	r = rw
	w = rw

	var rr MyReader
	var ww MyWriter
	rr = rw
	ww = rw
	fmt.Printf("\n%T - %T - %T\n",rw,rr,ww)

	//--------------------------------------------------
	var rwn MyNewReadWriter = os.Stdout		 //*os.File
	rn := rwn.(MyReader)
	wn := rwn.(MyWriter)
	fmt.Printf("\n%T - %T - %T\n",rwn,rn,wn)
	rn = rwn
	wn = rwn
	//rn = wn

	var rrn MyReader
	var wwn MyWriter
	rrn = rwn
	wwn = rwn
	fmt.Printf("\n%T - %T - %T\n",rwn,rrn,wwn)



}

func init(){
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"
	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"
}