package main

import (
	"fmt"
)

//N é definido o tamanho do vetor
//

func solucao(n int, s []int, d int, m int) int {
	resultado := 0
	for i := 0; i < n; i++ {
		//Definindo um limitador, caso ultrapasse o tamanho do vetor
		var limite int = i + m //Limite é definido pelo i do for + m do mês

		if limite > n {
			break //Se o limite for maior do que o n, acontecerá um loop
		}
		var soma int //Declaramos a variável sum para ser comparada ao dia.
		for j := i; j < limite; j++ {
			soma += s[j]
		}
		if soma == d {
			resultado++
		}

	}
	return resultado
}
func main() {
	var n, d, m int
	scan := fmt.Scan
	print := fmt.Print

	scan(&n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&s[i])
	}

	scan(&d, &m)
	resultado := solucao(n, s, d, m)

	print(resultado)

}
