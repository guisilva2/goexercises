package main

import "fmt"

const (
	//A constante BIG é definida e utilizando o prefixo '<<' alteramos a posição binária do número.
	//Por exemplo o número um será levado a esquerda 100 vezes
	Big   = 1 << 100
	Small = Big >> 99
	//Após isso ele retornará 99 vezes e retornará ao valor binário 10, que é o valor 2
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
