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
	array4 := make([]int, 5) //Array inicialized with size 5 and 0 for all values.
	array5 := make([][]int, 5)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
}
