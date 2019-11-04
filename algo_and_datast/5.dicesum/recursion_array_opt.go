package main

import (
	"fmt"



)




func printBinary(s int, a []int, sumSoFar int, sumint int) {
	//fmt.Println("s=", s, ", prefix=", prefix)

	if (s == 0) {

		sum := 0
		for _, num := range a {
			sum += num

		}

		//if sum == sumint {

			fmt.Print("{")
			for j := 0; j<len(a); j++ {
			
				fmt.Print(a[j])
				if j != len(a)-1 {
					fmt.Print(", ")
				}
			}
			fmt.Print("}\n")
			a = []int{}

		//}

		
		


		
			
	} else {

		
		for i := 1; i<7; i++ {
			
			if (sumSoFar + i + 1*(s-1) <= sumint) && (sumSoFar + i + 6*(s-1) >= sumint) {
				printBinary(s-1, append(a, i), sumSoFar+i, sumint)
			}
			

			

		}
			
	}
	

}








func main() {

	fmt.Println("Input the number")
	var inpint int
	fmt.Scanln(&inpint)

	var sumint int
	fmt.Scanln(&sumint)
	
	a := []int{}
	printBinary(inpint, a, 0,  sumint)

}



