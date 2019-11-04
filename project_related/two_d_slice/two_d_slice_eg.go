package main

import (
	"fmt"
)

func main() {
	var MultiDimensionalArray [3][5]int

	MultiDimensionalArray[2][1] = 5
	MultiDimensionalArray[1][1] = 2

	fmt.Println(MultiDimensionalArray)

}
