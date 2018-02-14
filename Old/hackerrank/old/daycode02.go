package main

import (
	//"bufio"
	"fmt"
	//"os"
	"strconv"
)

var _ = strconv.Itoa
var i uint64 = 4
var d float64 = 4.0
var s string = "HackerRank "
var a int = 1
var b float64 = 2
var c string = "ABC"

func main() {
	fmt.Println("Digite o valor inteiro de A:")
	fmt.Scan(&a)

	fmt.Println("Digite o valor inteiro de B:")
	fmt.Scan(&b)

	fmt.Println("Digite o valor inteiro de C:")
	fmt.Scan(&c)

	fmt.Printf("O Valor de A é : %d\n", a)
	fmt.Printf("O Valor de B é : %g\n", b)
	fmt.Println("O Valor de C é :" + c)

}

func imprimeVariaveis() {
	fmt.Println(i)
	fmt.Println(d)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
