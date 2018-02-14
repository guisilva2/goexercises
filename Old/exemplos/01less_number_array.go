package main

import "fmt"

func main() {
	number := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var less int = number[0]
	for i := 0; i < len(number); i++ {
		if number[i] < less {
			fmt.Println("O número", number[i], "é menor que o número", less)
			less = number[i]
		}
	}
	fmt.Println("O menor número encontrado foi o número", less)
}
