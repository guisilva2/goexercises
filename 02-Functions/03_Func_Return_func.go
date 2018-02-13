package main

import (
	"fmt"
)

func returnFunction() func() string {
	v1 := "ABC"
	return func() string {
		v1 = "DEF"
		return v1
	}
}

func main() {
	argFunc := returnFunction()
	fmt.Println(argFunc)
}
