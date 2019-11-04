package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("Foo:")
	//c2 := incrementor("Bar:")
	fmt.Println("Launching puller")
	//fmt.Println(c1)
	c3 := puller(c1)
	fmt.Println("Called puller")
	//c4 := puller(c2)
	//fmt.Println("Final Counter:", <-c3+<-c4)
	fmt.Println("Final Counter:", <-c3)
}

func incrementor(s string) chan int {
	fmt.Println("Inside incrementer")
	a := 10
	out := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			out <- 1
			fmt.Println(s, i)
			a = 11
		}
		close(out)
	}()
	fmt.Println(a)
	return out
}

func puller(c chan int) chan int {
	fmt.Println("Inside puller")
	out := make(chan int)
	b := 20
	go func() {
		var sum int
		for n := range c {
			sum += n
			//fmt.Printf("n: %d\n", n)
			fmt.Printf("sum: %d\n", sum)
			b = 21
		}
		out <- sum
		close(out)
	}()
	fmt.Println(b)
	return out
}
