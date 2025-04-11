package main

import ("fmt")

var opção int
var saldo int


 func deposito() { 
	 var deposito int 
	 fmt.Println("digite o valor do deposito")
	 fmt.Scan(&deposito)
	 saldo = saldo + deposito
	fmt.Println("seu saldo atual é:", saldo)
 }

 func saque() {
	var saque int
	fmt.Println("digite o valor do saque")
	fmt.Scan(&saque)
	saldo = saldo - saque
	fmt.Println("seu saldo atual é:", saldo)
 }
 
func main(){
fmt.Println("digite seu saldo atual")
fmt.Scan(&saldo)
 fmt.Println("digite 1 para depositar e 2 para saque")
 fmt.Scan(&opção)

 if opção == 1 {
	 deposito()
 } else if opção == 2 {
	saque()
 } 
 
}