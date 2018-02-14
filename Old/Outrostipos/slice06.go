package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	//Imprimindo o vetor S, o tamanho e a capacidade.
	if s == nil { //Se s for nulo irá imprimir na tela "nil!"
		fmt.Println("nil!")
	}
}
