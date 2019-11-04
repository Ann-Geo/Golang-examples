package main

import (
	"fmt"
)


func withPointer(x *int) {
	*x = *x * *x

}



type complexno struct {
	x, y int	

}



func newComplex(x, y int) *complexno {
	return &complexno{x, y}

}



func main() {

	x := -2
	withPointer(&x)
	fmt.Println(x)



	w := newComplex(x, -5)
	fmt.Println(*w)
	fmt.Println(w)

}
