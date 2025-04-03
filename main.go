package main

import "fmt"

func main() {
	var ages = [4]int{17, 16, 20, 40}
	nomes := [4]string{"mario", "luigi", "deadpool", "superman"}
	fmt.Println(ages)
    fmt.Println(nomes)
	nomes[3] = "clark kent"
	fmt.Println(nomes) 
		// Slice 
		var score = []int{100, 200, 300, 400}
		fmt.Println(score)
		score[1] = 2
		fmt.Println(score)
		fmt.Println(score, cap(score))
		rangeOne := score[1:3]
		fmt.Println(rangeOne)
		rangeTwo := score[2:]
		fmt.Println(rangeTwo)
		rangethree := score[:3]
		fmt.Println(rangethree)
		var superherois = []string{"deadpool", "homemaranha", "motoquerio fantasma"}
		fmt.Println(superherois)
		superherois = append(superherois, "ben10", "hulk")
		fmt.Println(superherois)
		fmt.Println(superherois, len(superherois), cap(superherois))
	
	}
