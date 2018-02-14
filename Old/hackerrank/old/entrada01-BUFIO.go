package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//A biblioteca bufio armazena o necessário para "input/output"
	leitor := bufio.NewReader(os.Stdin) //A biblioteca os, utiliza o sistema operacional para leitura, saída e erros de arquivo.
	frase, _ := leitor.ReadString('\n')

	fmt.Println("Foi digitado: " + frase)

}
