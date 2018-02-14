package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		//Default option
		fmt.Printf("%s.", os)
	}
}
