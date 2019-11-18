package main

import (
	"fmt"
)

func main() {
	//var a = []int{4, 10, 3, 5, 1}
	var a = []int{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17}
	sorted := heapSort(a, len(a))
	fmt.Println(sorted)
}

func heapSort(a []int, n int) []int {
	startIndex := n/2 - 1
	//build heap
	for i := startIndex; i >= 0; i-- {
		heapify(&a, n, i)
	}

	//one by one extract elements from heap
	for i := n - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(&a, i, 0)
	}

	return a

}

func heapify(a *[]int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if (l < n) && ((*a)[l] > (*a)[largest]) {
		largest = l
	}

	if (r < n) && ((*a)[r] > (*a)[largest]) {
		largest = r
	}

	if largest != i {
		(*a)[i], (*a)[largest] = (*a)[largest], (*a)[i]
		heapify(a, n, largest)
	}

}
