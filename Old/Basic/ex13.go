package main

import "fmt"

func main() {
	fmt.Printf("Podemos declarar uma variável sem definir o seu tipo, usando a sintaxe de declaração :=")
	a := 1
	b := 2.33
	c := "ola"
	d := "a"

	fmt.Printf("\nOs tipos das variáveis em sequência são %T, %T, %T, %T.", a, b, c, d)
}
