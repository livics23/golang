
package main

import ("fmt")
func dadosPessoa(idade int) (int, string){
	var condicao string
if idade > 18 {
	condicao = "maior de idade"

} else {
	condicao = "menor de idade"
} 
return idade, condicao
}

func main(){
var idade int 
 idade, condicao:= dadosPessoa(20)
fmt.Println(idade, condicao)
}