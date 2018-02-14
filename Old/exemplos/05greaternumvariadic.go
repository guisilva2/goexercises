package main

import "fmt"

func greaterNum(num ...int) int {
	greater := 0
	for _, n := range num {
		if n > greater {
			greater = n
		}
	}
	return greater
}

func main() {
	greater := greaterNum(2, 4, 6, 8, 12, 10, 2, 4, 6, 8, 12, 10, 27, 52, 122, 2)
	fmt.Println("O maior número é o:", greater)
}
