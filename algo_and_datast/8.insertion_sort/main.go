//insertion sort

package main


import "fmt"

func main() {


list_no := []int{4, 9, 5, 1, 0}






for i:=0;i<len(list_no);i++ {

	fmt.Printf("%d\n", i)

	for j:=i;j>=1;j-- {

	fmt.Printf("%d\n", j)

	fmt.Printf("%d, %d\n", list_no[j], list_no[j-1])

		if list_no[j]<list_no[j-1] {

			swap := list_no[j]

			list_no[j] = list_no[j-1]
			list_no[j-1] = swap


		}
		fmt.Println(list_no)
	}

	

	fmt.Println("****************")

}


fmt.Println(list_no)


}
