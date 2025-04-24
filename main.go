package main

import (
    "fmt"
)
func analisarNotas (nota1, nota2 float64) (float64, string) {
	 media := (nota1 + nota2) / 2 
	if media >= 6 {
		return media, "aprovado"
	}else{
	return media, "reprovado"
	
	}
    
    }
    func main() {
    media, resultado := analisarNotas (7.5, 5.5)
    fmt.Println("Média:", media)
    fmt.Println("Resultado:", resultado)
    }