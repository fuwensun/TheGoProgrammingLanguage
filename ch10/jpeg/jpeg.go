package main

import (
	"io"
	"image"
	"fmt"
	"os"
	"image/jpeg"
	"image/png"
)

//	***注意！！！***
//	./jpeg.exe <./png.png > jpg.jpg
//	cat ./png.png | ./jpeg.exe  > jpg1.jpg

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil{
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error{
	img, kind, err := image.Decode(in)
	if err != nil{
		fmt.Println(png.BestCompression)
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality:95})

}