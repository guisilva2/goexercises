package main

import (
	"fmt"
)

func sumNumbers(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func main() {

	fmt.Println(sumNumbers(1, 2, 3, 4, 5))
}
