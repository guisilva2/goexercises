package main

import (
	"fmt"
)

func main() {
	var n int
	var simb string

	fmt.Println("*** Árvore de #. ***")
	fmt.Println("Tamanho da árvore: ")
	fmt.Scan(&n)

	simb = "#"

	var j int
	for i := 0; i < n; i++ {
		for j = n - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}
		fmt.Println(simb)
		simb += "#"
	}

}
