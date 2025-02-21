package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a Golang program!")
	fmt.Println("Golang version:", runtime.Version())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)

}
