//quick sort

package main


import (

"fmt"
//"math/rand"
//"time"

)


func main() {

	nums := []int{2, 6, 5, 3, 8, 7, 1, 0}
	//nums := []int{0}


	fmt.Println(quickSort(nums))


}

func quickSort(nums []int) []int {

	//numsCp := make([]int, len(nums))

	//copy(numsCp, nums)

	//fmt.Println(numsCp)

	quickSortHelper(&nums, 0, len(nums)-1)


	return nums

}



func quickSortHelper(nums *[]int, left, right int){


if left>=right {
return
}

pivot := (*nums)[(left+right)/2]

partPoint := partition(nums, left, right, pivot)

quickSortHelper(nums, left, partPoint-1)
quickSortHelper(nums, partPoint, right)


}



func partition(nums *[]int, left, right, pivot int) int{	

	//pivotIndex := left + rand.Int() % (right-left)
	//pivotIndex := (*nums)[(left+right)/2]
	//fmt.Println("Pivot index", pivotIndex)
	//(*nums)[pivotIndex], (*nums)[right] = (*nums)[right], (*nums)[pivotIndex]

	for left <= right{

		for ((*nums)[left] < pivot) {
			left++

		}

		for ((*nums)[right] > pivot) {
			right--

		}

		if (left <= right) {

			temp := (*nums)[left]
			(*nums)[left] = (*nums)[right]
			(*nums)[right] = temp

			left++
			right--
		}
	}


	return left
}

