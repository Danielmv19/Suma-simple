package main

import "fmt"

func main() {
	//primer numero ingresado
	fmt.Println("numero 1: ")
	var a int
	fmt.Scanln(&a)
	//segundo numero ingresado
	fmt.Println("numero 2: ")
	var b int
	fmt.Scanln(&b)
	//suma
	suma := a + b
	fmt.Println("la suma es: ")
	fmt.Println(suma)

}
