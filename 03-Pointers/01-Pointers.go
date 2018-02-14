package main

import "fmt"

//Work's
func swap(number1, number2 *int) {
	*number1, *number2 = *number2, *number1
}

//Don't work
func swap2(number1, number2 *int) {
	*number1 = *number2
	*number2 = *number1
}

func rotate(args ...*int) {
	firstValue := *args[0]
	for i := 0; i < len(args)-1; i++ {
		*args[i] = *args[i+1]
	}
	*args[len(args)-1] = firstValue
}

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5

	rotate(&a, &b, &c, &d, &e)
	fmt.Println(a, b, c, d, e)
}
