package main

import (
	"fmt"
)



func SumandDiff(a, b int) (sum, diff int) {

	sum = a + b
	diff = a - b

	return


}

func main() {

	fmt.Println(SumandDiff(5, 7))

}
