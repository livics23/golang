package main
import "fmt"
 
 func main() {
	var saldo int
	var digite int 
	var deposito int
	var saque int
	var saquefinal int
	var depositofinal int


fmt.Println("digite seu saldo")
fmt.Scan(&saldo)
fmt.Println("digite 1 para depositar e 2 para saque")
fmt.Scan(&digite)


if digite == 1{
	fmt.Println("coloque o valor do seu deposito")
	fmt.Scan(&deposito)
	depositofinal = saldo + deposito
	fmt.Println("seu saldo final é :", depositofinal)
} else if digite == 2 {
	fmt.Println("coloque o valor que deseja retirar")
	fmt.Scan(&saque)
	saquefinal = saldo - saque 
	fmt.Println("seu saldo final é :", saquefinal)
} 
 }