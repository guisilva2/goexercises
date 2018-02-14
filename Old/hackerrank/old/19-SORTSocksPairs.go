package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, pares int
	scan := fmt.Scan
	print := fmt.Println

	scan(&n)
	socks := make([]int, n)
	//Leitura das "cores das meias
	for i := 0; i < n; i++ {
		scan(&socks[i])
	}
	sort.Ints(socks)
	//Percorremos o vetor do início ao fim
	for i := 0; i < n-1; i++ {
		if socks[i] == socks[i+1] { //Se o valor for igual ao próximo
			pares++ //Adicionamos um par à variável
			i++     //O i é incrementado 2 vezes, pois passamos o próximo número que já foi comparado e acessamos o próximo.
		}
	}
	print(pares)
}
