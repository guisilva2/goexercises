package main

import (
	"fmt"
)

func main() {

	var mealCost float64
	var tipPercent int
	var taxPercent int
	var tip float64
	var tax float64
	var totalCost int

	leitor.Scan() //O leitor é acionado através da função Scan, e pronto para leitura dos dados
	fmt.Scan(&mealCost)

	fmt.Scan(&tipPercent)

	fmt.Scan(&taxPercent)

	tip = mealCost * (float64(tipPercent) / 100)
	tax = mealCost * (float64(taxPercent) / 100)
	totalCost = Round(mealCost + tip + tax)
	fmt.Println()
	fmt.Println(tip)
	fmt.Println(tax)
	fmt.Printf("The total meal cost is %d dollars.", totalCost)

}
