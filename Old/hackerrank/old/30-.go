package main

import (
	"fmt"
)

func main() {
	var n int
	scan := fmt.Scan
	print := fmt.Println

	scan(&n)

	var array [5]int

	for i := 0; i < len(array); i++ {
		fmt.Scan(&array[i])
		array[i] = n % 10
		n /= 10
	}

	print(array)

}
