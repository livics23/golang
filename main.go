
package main

import ("fmt")
func main(){
	alunoIdade := make(map[string]int)
	alunoIdade ["livia"] = 16 
	alunoIdade ["laura"] = 15
	fmt.Println("idade livia", alunoIdade["livia"])
notasAluno := map[string]float64{
	"livia" :1.6,
     "kelly": 8.5,
}
for nome, nota := range notasAluno{
	fmt.Println("%s tirou a nota %.1f \n", nome, nota,)
}
}