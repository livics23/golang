package main

import (
	"fmt"
	"strings"
	"sort"
	
)
func main() {
	greeting := "hello my friends!"
	fmt.Println(strings.Contains(greeting, "friends"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "my"))
	fmt.Println(strings.Split(greeting, "!"))
	ages := []int {90, 80, 900}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 80)
	fmt.Println(index)
	names := []string{"carol", "maicon", "diegoike"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names,"diegoike"))
	x := 0
	
	for x < 5 {
		fmt.Println(x)
		x++
	}
	
	for i :=0; i >5; i++{
		fmt.Println("for 2: ", i)
	}
	for i:=0; i <len(names); i++{
		fmt.Println(names[i])
	}
	for index, value := range names{
		fmt.Println("índice é", index, "e o valor", value,)
	}

	for index, value := range ages{
		fmt.Println("o indice é", index, "e o valor", value)

	}

	superheois := []string{"deadpool", }
	
	
}
	