package main

import (
	"fmt"
)

func main() {
	var n, sumcol1, sumcol2 int

	fmt.Println("Por favor digite o número de linhas / colunas da matrix:")
	fmt.Scanf("%d", &n)

	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, n)
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			fmt.Printf("Posição [%d], [%d]\n", x+1, y+1)
			fmt.Scanf("%d", &matriz[x][y])
		}
	}
	b := 0
	for a := 0; a < n; a++ {

		sumcol1 += matriz[a][b]
		b++
	}
	b = n - 1
	for a := 0; a < n; a++ {

		sumcol2 += matriz[a][b]
		b--
	}

	res := (sumcol1 - sumcol2)

	if res < 0 {
		res *= -1
	}

	fmt.Println((res))

}
