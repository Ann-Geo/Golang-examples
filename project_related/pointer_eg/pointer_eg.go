package main

import "fmt"

func main() {
	a := 1
	p := &a
	fmt.Println(a, p)
	fmt.Printf("%T, %T\n", a, p)
}
