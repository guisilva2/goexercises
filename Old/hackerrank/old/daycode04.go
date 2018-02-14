package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)
	if n%2 == 1 {
		fmt.Print("Weird")
	} else if (n%2 == 0) && (n > 2) && (n < 5) {
		fmt.Print("Not Weird")
	} else if (n%2 == 0) && (n > 6) && (n < 20) {
		fmt.Print("Weird")
	} else if (n%2 == 0) && (n > 20) {
		fmt.Print("Not Weird")
	}
}
