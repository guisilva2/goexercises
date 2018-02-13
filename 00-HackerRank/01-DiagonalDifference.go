package main

import (
	"fmt"
	"math"
)

func main() {
	var n, diag1, diag2 int
	fmt.Scan(&n)
	array := make([][]int, n)
	for i := range array {
		array[i] = make([]int, n)
	}
	for i := range array {
		for j := range array {
			fmt.Scan(&array[i][j])
			if i == j {
				diag1 += array[i][j]
			}
			if i+j == n-1 {
				diag2 += array[i][j]
			}
		}
	}
	result := math.Abs(float64(diag1) - float64(diag2)) //Or create a function Abs returning -number
	fmt.Print(result)
}
