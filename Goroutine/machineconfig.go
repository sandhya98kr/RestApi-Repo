package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("logical processors:", runtime.NumCPU())
	fmt.Println("Operating system:", runtime.GOOS)
	fmt.Println("System architecture:", runtime.GOARCH)
	fmt.Println("Maximum processes", runtime.GOMAXPROCS(5))
}
