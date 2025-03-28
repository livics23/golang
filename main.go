package main

import "fmt"

func main() {
	var usuario string
	var senha string
	fmt.Println("digite a senha")
	fmt.Scan(&senha) 
	fmt.Println("digite o nome de usuario")
	fmt.Scan(&usuario)
	if senha == "1234" && usuario == "admin" { 
		fmt.Println("Acesso permitido")
	} else { 
		fmt.Println("Acesso negado")
	}
}
