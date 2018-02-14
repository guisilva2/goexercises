package main

import "fmt"

func squarePointer(x *float64) {
	*x = *x * *x
}

func multiply(n int, num *float64) {
	for i := 0; i < n; i++ {
		*num = *num * *num
		fmt.Printf("Potência n %d = %.2f \n", i+1, *num)
	}
}

func main() {
	x := 1.5
	squarePointer(&x)
	fmt.Println("Square ===>", x)
	multiply(5, &x)
	fmt.Printf("Square ===> %.2f", x)

}
