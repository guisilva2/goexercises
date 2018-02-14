package main

import "fmt"

func main() {
	fmt.Println("Exemplos básicos utilizando maps.")
	contas := map[int]string{
		1: "Guilherme",
		2: "Romário",
		3: "Robson",
		4: "Hudson",
	}

	fmt.Println("Conta:", contas[1])
	fmt.Println(contas)

}
