//bubble sort

package main


import "fmt"

func main() {


list_no := []int{4, 9, 5, 1, 0}


l := len(list_no)
l = l-1


for i:=0;i<len(list_no)-1;i++ {

	fmt.Println(l)
	for j:=0;j<l;j++ {

	fmt.Printf("%d, %d\n", list_no[j], list_no[j+1])

		if list_no[j]>list_no[j+1] {

			swap := list_no[j]

			list_no[j] = list_no[j+1]
			list_no[j+1] = swap


		}
	}

	fmt.Println(list_no)

	l = l-1


	fmt.Println(l)
	fmt.Println("****************")

}


fmt.Println(list_no)


}
