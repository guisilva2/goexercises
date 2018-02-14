package main

import "fmt"

var maxvelocidade int = 100
var minvelocidade int = 0
var peso float32 = 4079
var ligado bool = false

func imprimeVariaveis() {
	fmt.Println(maxvelocidade)
	fmt.Println(minvelocidade)
	fmt.Println(peso)
	fmt.Println(ligado)
}

func main() {

	imprimeVariaveis()

}
