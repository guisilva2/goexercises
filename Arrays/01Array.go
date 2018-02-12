package main

import (
	"fmt"
)

func main() {
	var array1 = []int{
		1,
		2,
		3,
		4,
		5,
	}
	array2 := []int{1, 2, 3, 4, 5}
	array3 := [][]int{
		{1, 2},
		{2, 3},
		{4, 5},
	}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}
