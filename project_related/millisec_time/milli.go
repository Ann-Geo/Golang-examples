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

	//t1 := time.Now()
t1 := time.Now()

fmt.Printf("%T\n", t1)

fmt.Println(t1.Format(time.StampMilli))

time.Sleep(33*time.Millisecond)

	t2 := time.Now()
fmt.Println(t2.Format(time.StampMilli))

duration := t2.Sub(t1)
fmt.Println(duration.Seconds())

}
