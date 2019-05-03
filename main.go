package main

import (
	"fmt"

	mypkg "github.com/mr-dm/go_simple_example_mypkg" // this is now imported as mypkg
)

func main() {
	fmt.Printf("Hello, world. Hi %s\n", mypkg.AddPrefix("Go语言"))
}
