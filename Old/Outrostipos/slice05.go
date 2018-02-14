package main

import "fmt"

//Função que imprime a Slice, o literal %d define um número decimal
func printSlice(s []int) {
	fmt.Printf("len%d cap%d %v\n", len(s), cap(s), s)
	//LENGHT ==> Recebe o tamanho da matriz, valores que estão sendo ocupados, len, CAP ==> Capacidade da matriz
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
}
