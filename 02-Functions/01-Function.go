package main

import "fmt"

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func calculator(num1 int, operator string, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		fmt.Println("Unknow operation: We don't calculate this for u :)", operator)
	}
	return 0
}

func main() {
	var n1, n2 int
	var op string
	fmt.Println(sum(5, 2))
	fmt.Println("*****CALCULATOR*****\nInsert the first number")
	fmt.Scan(&n1)
	fmt.Println("Operation: + - / * ")
	fmt.Scan(&op)
	fmt.Println("Insert the second number")
	fmt.Scan(&n2)
	fmt.Println(calculator(n1, op, n2))

}
