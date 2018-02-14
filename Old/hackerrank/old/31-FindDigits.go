package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nTest int

	fmt.Scan(&nTest)
	num := make([]int, nTest)
	for i := 0; i < nTest; i++ {
		fmt.Scan(&num[i])
	}
	intStr := strconv.Itoa(num[0])
	intStr2 := []byte(num[0])
	fmt.Println(intStr2)

}
