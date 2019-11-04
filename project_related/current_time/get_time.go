package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%T\n",t)
	fmt.Println(t)
	t1 := t.Format(time.RFC850)
	fmt.Printf("time in %T format - %v\n", t1, t1)
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())

}
