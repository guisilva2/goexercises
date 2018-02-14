package main

import "fmt"

func primeiro() {
	fmt.Println("Primeiro")
}

func segundo() {
	fmt.Println("Segundo")
}

func main() {
	defer segundo() //Defer adia a execução, e é executada após a função atual.
	primeiro()
	primeiro()
	fmt.Println("A função com o atributo defer será inicializada após ao término da função main.")
}
