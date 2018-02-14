package main

import (
	"fmt"
)

func main() {
	var n, k, maior int
	scan := fmt.Scan
	print := fmt.Println
	scan(&n, &k)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&vetor[i])
		if vetor[i] > k && vetor[i] > maior {
			maior = vetor[i]
		}
	}
	if maior != 0 {
		print(maior - k)
	} else {
		print(0)
	}

}
