package main

import "fmt"

func main() {
	array1 := [][]int{}

	//IINITIALIZE ARRAY WITH 0's values ==> TYPE ==> SIZE
	array2 := make([][]int, 3)
	for i := range array1 {
		array1[i] = make([]int, 3)
	}
	fmt.Println(array1, array2)
}
