package main

import (
	"fmt"
	"strings"
)

func main() {

	// Declarando um array do tipo inteiro com os valores
	texts := []string{"A", "B", "C"}
	valuesText := []string{} //Os valores irão para o vetor do tipo string

	// strconv.Itoa. ==> Itoa converte int para string
	for i := range texts {
		text := texts[i]
		//	text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	// O método Join irá, acresscentar o elemento no final do vetor ou array.
	result := strings.Join(valuesText, "+")
	fmt.Println(result)
}

/*
for i := 0; i < len(baralho); i++ {
		for j := 0; j < len(baralho); j++ {
			if baralho[i].valor > baralho[j].valor { //O valor armazenado na posição i é comparado, se for maior:
				aux[0] = baralho[i]     //O auxiliar recebe o valor armazenado na posição j
				baralho[i] = baralho[j] //O valor da posição j é jogado na posição i
				baralho[j] = aux[0]     //E o valor da posição i recebe o auxiliar
			} //Ou seja é feita uma troca de valores, para ser realizada a ordenação
		}
	}
*/
