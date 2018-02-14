package main

import "fmt"

func main() {
	//A função make obtém a matriz e retorna como uma slice.
	pow := make([]int, 10)
	//A função range acessa os valroes da matriz.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	//O for inicia a partir de 0, após isso, o valor 1 é acrescentado de casas decimais à direita.
	//Por exemplo 10, após ==> 100, 1000, 100000, 100000 ==> 1,2,4,8,16,32,64 ...
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
