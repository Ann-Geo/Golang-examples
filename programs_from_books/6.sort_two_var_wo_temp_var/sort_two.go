package main

import (
	"fmt"
)



func sortTwo(a, b int) (int, int) {
	if a > b {
		return b, a
	}else {
		return a, b

	}
	
	

}





func main() {

	fmt.Println(sortTwo(4, 1))




}
