package main

import "fmt"

func main() {
	//var a = []int{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17}
	var a = []int{4, 10, 3, 5, 1}
	heap := buildHeap(a, len(a))
	fmt.Println(heap)
}

func buildHeap(a []int, n int) []int {
	startIndex := n/2 - 1
	for i := startIndex; i >= 0; i-- {
		heapify(&a, len(a), i)
	}

	return a
}

func heapify(a *[]int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	//left child larger than root
	if (l < n) && (*a)[l] > (*a)[largest] {
		largest = l
	}
	//right child larger than root
	if (r < n) && (*a)[r] > (*a)[largest] {
		largest = r
	}
	//largest is not root
	if largest != i {
		temp := (*a)[i]
		(*a)[i] = (*a)[largest]
		(*a)[largest] = temp
		heapify(a, n, largest)
	}
}
