package main

import (
	"fmt"
)

func soma(num1 int, num2 int) int {
	return num1 + num2
}
func sum(n1, n2 int) int {
	return n1 + n2
}
func main() {
	fmt.Println("Iniciando com funções:")
	fmt.Println("Exemplo de função, utilizando a soma de 2 números")
	fmt.Println(soma(20, 20))

	fmt.Println("A segunda função tem a declaração de váriaveis simplificada:")
	fmt.Println(sum(20, 30))
}
