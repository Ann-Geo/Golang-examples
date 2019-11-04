package main

import (
	"fmt"


)




func printBinary(s int, prefix string) {
	//fmt.Println("s=", s, ", prefix=", prefix)

	if (s == 0) {
		fmt.Println(prefix)
	} else {

		
		printBinary(s-1, prefix+"0")
		printBinary(s-1, prefix+"1")
		printBinary(s-1, prefix+"2")
		printBinary(s-1, prefix+"3")
		printBinary(s-1, prefix+"4")
		printBinary(s-1, prefix+"5")
		printBinary(s-1, prefix+"6")
		printBinary(s-1, prefix+"7")
		printBinary(s-1, prefix+"8")
		printBinary(s-1, prefix+"9")


	
	}
	

}




func main() {

	fmt.Println("Input the number")
	var inpint int
	fmt.Scanln(&inpint)
	
	
	printBinary(inpint, "")

}



