package main

import (
	"fmt"
)



func NextIntFunc() func() int {

	i := 0
	return func() int {
		i++
		return i

	}




}

func main() {

	nextInt := NextIntFunc()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	nextnextInt := NextIntFunc()

	fmt.Println(nextInt())

	fmt.Println(nextnextInt())




}
