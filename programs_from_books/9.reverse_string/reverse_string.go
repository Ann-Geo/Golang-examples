package main

import (
	"fmt"

)




func reverseString(s string) string {
	//outp := []rune(inpstr)
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)


}




func main() {

	fmt.Println("Input the string")
	var inpstr string
	fmt.Scanln(&inpstr)
	
	fmt.Println(inpstr)
	fmt.Println(string(inpstr[1]))
	fmt.Println(reverseString(inpstr))

}



