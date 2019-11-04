package main

import "fmt"
import "strings"

func main() {
	x := "Tuesday, 16-Apr-19 11:02:03 EDT"

	i := strings.Index(x, "EDT")
	fmt.Println("Index: ", i)
	if i > -1 {
		chars := x[i-3 : i]
		fmt.Printf("%T\n", chars)
		arefun := x[i+1:]
		fmt.Println(chars)
		fmt.Println(arefun)
	} else {
		fmt.Println("Index not found")
		fmt.Println(x)
	}
}
