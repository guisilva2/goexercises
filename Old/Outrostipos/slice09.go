package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	//A função append funciona em matrizes vazias.
	s = append(s, 0)
	printSlice(s)

	//A slice ou matriz, cresce conforme necessário.
	s = append(s, 1)
	printSlice(s)

	//Podemos adicionar um ou mais elementos com a função append.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
