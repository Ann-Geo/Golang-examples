package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("/home/research/goworkspace/src/vsc_workspace/images/image0000000.png") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%T\n", b) // print the content as 'bytes'

	//str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'
}
