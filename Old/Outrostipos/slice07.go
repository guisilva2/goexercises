package main

import "fmt"

func main() {
	//Com a função make podemos receber a matriz e retorná-la como uma slice.
	a := make([]int, 5)
	printSlice("a", a) //Acessando a função printSlice, será impresso a matriz, o leght e a capacidade.

	b := make([]int, 5, 10)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
