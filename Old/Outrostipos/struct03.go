package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  //Irá receber o valor diretamente
	v2 = Vertex{X: 1}  //Deixará o X implícito
	v3 = Vertex{}      //Exibirá is valores como 0
	p  = &Vertex{1, 2} //Irá obter o valor de Vertex e armazená-lo como ponteiro
)

func main() {
	fmt.Println()
	{
		fmt.Println(v1, p, v2, v3)
	}

}
