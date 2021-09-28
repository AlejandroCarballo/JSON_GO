package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Fprintln(os.Stdout, "Hello Ale")
	io.WriteString(os.Stdout, "Hello, Ale")

}
