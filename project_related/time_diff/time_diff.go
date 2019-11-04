package main

import (
	"fmt"
	"time"
)

//const (
// See http://golang.org/pkg/time/#Parse
//timeFormat = "Monday, 13-May-19 14:08:42 EDT"
//)

func main() {
for {
	v := time.Now().Format(time.RFC850)
	then, err := time.Parse(time.RFC850, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	duration := time.Since(then)
	fmt.Println(duration.Seconds())
}
/*
	v = time.Now().Format(time.RFC850)
	then, err = time.Parse(time.RFC850, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	duration = time.Since(then)
	fmt.Println(duration.Seconds())

	v = time.Now().Format(time.RFC850)
	then, err = time.Parse(time.RFC850, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	duration = time.Since(then)
	fmt.Println(duration.Seconds())

	v = time.Now().Format(time.RFC850)
	then, err = time.Parse(time.RFC850, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	duration = time.Since(then)
	fmt.Println(duration.Seconds())
*/
}
