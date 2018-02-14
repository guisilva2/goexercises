package main

import "fmt"

func swap(palavra1, palavra2 string) (string, string) {
	return palavra2, palavra1
}
func main() {
	fmt.Println("A função 'SWAP', faz a troca das palavras, recebendo duas strings e retornando as mesmas invertidas.")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
