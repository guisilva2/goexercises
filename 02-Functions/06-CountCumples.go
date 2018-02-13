package main

import (
	"fmt"
)

func countClumps(numbers []int) int {
	var count, pairs = 0, 1
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] == numbers[i+1] {
			pairs++
			if pairs == 2 {
				count++
			}
		} else {
			pairs = 1
		}
	}
	return count
}

func main() {
	numbers := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	fmt.Println(countClumps(numbers))
}
