package main

import (
	"fmt"
)

func main() {
	var n, positivos, negativos, zeros int

	fmt.Println("*** Conte os números positivos, negativos e os zeros. ***")
	fmt.Println("Tamanho do vetor / array: ")
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println("Posição:", i+1)
		fmt.Scan(&vetor[i])
		if vetor[i] < 0 {
			negativos++
		} else if vetor[i] > 0 {
			positivos++
		} else {
			zeros++
		}
	}

	var n1 float64 = float64(positivos) / float64(n)
	var n2 float64 = float64(negativos) / float64(n)
	var n3 float64 = float64(zeros) / float64(n)

	fmt.Printf("%.6f\n%.6f\n%.6f", n1, n2, n3)

}
