package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //Ponteiro para o I
	fmt.Println(*p) // Irá ler o I através do P
	*p = 21         // Setar o I, através do P
	fmt.Println(i)  // Observar o novo valor de I

	p = &j         // Ponteiro para o J
	*p = *p / 37   //Irá dividir o J através do P, ou seja estamos acessando o J através do P
	fmt.Println(j) // Observe o novo valor de J
}
