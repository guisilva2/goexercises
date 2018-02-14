package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	//A propriedade range, percorre os valores da matriz.
	//Por exemplo, acessa os valores, no caso começando à partir da posição 0 da matriz
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
