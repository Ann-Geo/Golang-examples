package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	minusI := false
	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			minusI = true
			break
		}
		
	}

	if minusI == true {
		fmt.Println("found -i")
		fmt.Println("y/n:")
		var answer string
		fmt.Scanln(&answer)
		fmt.Println("You entered:", answer)
	}else {
		fmt.Println("-i not found")
	}

}
