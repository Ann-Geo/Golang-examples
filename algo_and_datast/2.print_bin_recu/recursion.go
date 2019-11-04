package main

import (
	"fmt"


)




func printBinary(s int, prefix string) {
	fmt.Println("s=", s, ", prefix=", prefix)

	if (s == 0) {
		fmt.Println(prefix)
	} else {
		printBinary(s-1, prefix+"0")
		printBinary(s-1, prefix+"1")
	
	}
	

}




func main() {

	fmt.Println("Input the number")
	var inpint int
	fmt.Scanln(&inpint)
	
	
	printBinary(inpint, "")

}



