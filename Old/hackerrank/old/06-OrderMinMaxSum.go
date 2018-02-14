package main

import "fmt"

func main() {

	vN := [5]int{}
	var aux, minSum, maxSum int
	for i := 0; i < 5; i++ {
		fmt.Scan(&vN[i])
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if vN[i] > vN[j] { //O valor armazenado na posição i é comparado, se for maior:
				aux = vN[j]   //O auxiliar recebe o valor armazenado na posição j
				vN[j] = vN[i] //O valor da posição j é jogado na posição i
				vN[i] = aux   //E o valor da posição i recebe o auxiliar
			} //Ou seja é feita uma troca de valores, para ser realizada a ordenação
		}
	}

	minSum = vN[4] + vN[3] + vN[2] + vN[1]
	maxSum = vN[0] + vN[1] + vN[2] + vN[3]

	fmt.Println(minSum, maxSum)

}
