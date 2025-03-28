package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	fmt.Print("digite um numero")
	fmt.Scan(&numero1)
	fmt.Println("digite outro numero")
	fmt.Scan(&numero2)

	fmt.Println(" a soma é ", numero1+numero2)
	fmt.Println("a subtração é", numero1-numero2)
	fmt.Println("a divisão é", numero1/numero2)
	fmt.Println("a multiplicação é", numero1*numero2)
	fmt.Println("o resto da divisão é", numero1%numero2)
}
