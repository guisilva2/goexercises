package main

import (
	"fmt"
)

func main() {
	var n, maior, maiorIndice int
	tipo := make([]int, 5)
	scan := fmt.Scan
	print := fmt.Println

	scan(&n)
	v1 := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&v1[i])
		switch v1[i] {
		case 1:
			tipo[0]++
		case 2:
			tipo[1]++
		case 3:
			tipo[2]++
		case 4:
			tipo[3]++
		case 5:
			tipo[4]++
		}
	}

	for i := 0; i < 5; i++ {
		if tipo[i] > maior {
			maior = tipo[i]
			maiorIndice = i
		}
	}

	print(maiorIndice + 1)

}
