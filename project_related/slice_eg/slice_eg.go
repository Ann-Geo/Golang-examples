package main

import (
	"fmt"
)

func main() {
	var Segment = make([]int, 3)
	fmt.Println(Segment, len(Segment), cap(Segment))
	Segment = nil
	fmt.Println(Segment, len(Segment), cap(Segment))
	Segment = append(Segment, 100)
	fmt.Println(Segment, len(Segment), cap(Segment))
}
