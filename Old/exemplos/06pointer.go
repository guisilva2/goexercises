package main

import "fmt"

func main() {
	var n int
	n = 5
	fmt.Println("A variável n tem o valor de:", n, "e está armazenada no endereço.", &n)
	fmt.Println("Valor ===>", n, " Endereço ===>", &n)
}
