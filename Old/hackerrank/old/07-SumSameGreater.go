package main

import "fmt"

func main() {

	var nVelas, qVelas, maior int
	//fmt.Println("Quantidade de velas:")
	fmt.Scan(&nVelas)

	vetor := make([]int, nVelas)

	for i := 0; i < nVelas; i++ {
		fmt.Scan(&vetor[i])

		if vetor[i] == maior {
			qVelas++
		}
		//Se o valor for maior que o anterior e maior que o maior armazenado é inserido na variável.
		if vetor[i] > maior {
			maior = vetor[i]
			qVelas = 1
		}
	}

	fmt.Println(qVelas)
}
