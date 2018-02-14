package main

import "fmt"

func main() {
	var a, b [3]int
	var alice, bob int
	// /	var c [2]int

	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &b[i])
	}

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice += 1
		} else if b[i] > a[i] {
			bob += 1
		}
	}

	fmt.Printf("Pontos de Alice: %d \n Pontos de Bob: %d", alice, bob)
}
