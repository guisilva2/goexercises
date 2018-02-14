package main

import (
	"fmt"
)

func main() {
	var n, aux, pares int
	scan := fmt.Scan
	print := fmt.Println

	scan(&n)
	socks := make([]int, n)
	//Leitura das "cores das meias
	for i := 0; i < n; i++ {
		scan(&socks[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if socks[i] > socks[j] { //O valor armazenado na posição i é comparado, se for maior:
				aux = socks[j]      //O auxiliar recebe o valor armazenado na posição j
				socks[j] = socks[i] //O valor da posição j é jogado na posição i
				socks[i] = aux      //E o valor da posição i recebe o auxiliar
			} //Ou seja é feita uma troca de valores, para ser realizada a ordenação
		}
	}
	//Percorremos o vetor do início ao fim
	for i := 0; i < n-1; i++ {
		if socks[i] == socks[i+1] { //Se o valor for igual ao próximo
			pares++ //Adicionamos um par à variável
			i++     //O i é incrementado 2 vezes, pois passamos o próximo número que já foi comparado e acessamos o próximo.
		}
	}
	print(pares)
}
