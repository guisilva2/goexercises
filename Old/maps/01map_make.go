package main

import "fmt"

func main() {
	fmt.Println("Exemplos básicos utilizando maps.")
	elementos := make(map[string]string)
	elementos["H"] = "Hidrogênio"
	elementos["O"] = "Oxigênio"
	elementos["H"] = "Hydrogen"
	elementos["He"] = "Helium"
	elementos["Li"] = "Lithium"
	elementos["Be"] = "Beryllium"
	elementos["B"] = "Boron"
	elementos["C"] = "Carbon"
	elementos["N"] = "Nitrogen"
	elementos["O"] = "Oxygen"
	elementos["F"] = "Fluorine"
	elementos["Ne"] = "Neon"

	fmt.Println(elementos)
	if elemento, ok := elementos["FGD"]; ok {
		fmt.Println(elemento, "presente.")
		fmt.Println(ok)
	} else {
		fmt.Println(elemento, "elemento não encontrado.")
	}
}
