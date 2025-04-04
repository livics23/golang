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
	
}
	