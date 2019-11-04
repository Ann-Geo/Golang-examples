package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	sum := 0
	for i := 0; i < len(arguments); i++ {
		input, _ := strconv.Atoi(arguments[i])
		sum = sum + input
	}

	fmt.Println(sum)

}
