package main

import (
	"fmt"
	"strconv"
)
func printBinary(s int, prefix string) {
	//fmt.Println("s=", s, ", prefix=", prefix)

	if (s == 0) {
		fmt.Println(prefix)
	} else {

		for i := 0; i<10; i++ {
		
			printBinary(s-1, prefix+strconv.Itoa(i))		
		}
	}
}
func main() {

	fmt.Println("Input the number")
	var inpint int
	fmt.Scanln(&inpint)
	
	
	printBinary(inpint, "")
}
