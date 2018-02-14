package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, sum int

	fmt.Scanf("%d", &n)
	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vetor[i])
		sum = vetor[i] + sum
	}
	fmt.Print(sum)
}
