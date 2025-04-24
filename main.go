package main

import "fmt"

func main(){
	estoque := map[string]int{


	"coxinha" :  10,
	"pao de queijo" : 15,
   "refrigerante" : 20,
}
	
	estoque["coxinha"] -=2
   estoque["pao de queijo"] -=1
	

	fmt.Println("estoque atual:")
	for produto , quantidade := range estoque{
		fmt.Printf("%s, %d \n", produto, quantidade) 
	}
}
