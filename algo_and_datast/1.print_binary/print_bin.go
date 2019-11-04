package main

import (
	"fmt"
	"math"

)




func printBinary(s float64) {
	limit := math.Pow(2, s)
	var i float64
	for i = 0; i < limit; i++ {
		fmt.Printf("%03b\n", int64(i))
		//fmt.Println(i)
	}


	


}




func main() {

	fmt.Println("Input the number")
	var inpint float64
	fmt.Scanln(&inpint)
	
	
	printBinary(inpint)

}



