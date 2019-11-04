//selection sort

package main


import "fmt"

func main() {


list_no := []int{4, 9, 5, 1, 0}




for i:=0;i<len(list_no)-1;i++ {
	for j:=i+1;j<len(list_no);j++ {

	fmt.Printf("%d, %d\n", i, j)

		if list_no[i]>list_no[j] {

			swap := list_no[j]

			list_no[j] = list_no[i]
			list_no[i] = swap


		}
	}

}


fmt.Println(list_no)


}
