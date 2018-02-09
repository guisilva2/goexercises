package main

import (
	"fmt"
)

func main() {
	aPoints, bPoints := 0, 0
	a, b := make([]int, 3), make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	for i, v := range a {
		if v > b[i] {
			aPoints++
		} else if v < b[i] {
			bPoints++
		}
	}
	fmt.Print(aPoints, bPoints)
}
