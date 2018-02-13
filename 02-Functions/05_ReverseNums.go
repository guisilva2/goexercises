package main

import (
	"fmt"
)

//UnnecessaryComment - UC - Reverse with with 1 variable on for loop
func reverse(v1 []int) []int {
	vetReversed := make([]int, len(v1))
	for i := range v1 {
		vetReversed[i] = v1[len(v1)-(i+1)]
	}
	return vetReversed
}

//UC - Reverse with with 2 variables on for loop
func reverse2(v1 []int) []int {
	vetReversed := make([]int, len(v1))
	for i, v := range v1 {
		vetReversed[len(v1)-(i+1)] = v
	}
	return vetReversed
}

func main() {
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reverse2(v1))
}
