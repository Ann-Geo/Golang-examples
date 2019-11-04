package main

import "fmt"

type ImageStore struct {
	a int
	b int
}

var ImageStoreLog []ImageStore

var LogIndex int = -1

func main() {

	m := make(map[int]*ImageStore)

	imStore1 := ImageStore{a: 1, b: 1}
	imStore2 := ImageStore{a: 2, b: 2}
	imStore3 := ImageStore{a: 3, b: 3}
	//imStore4 := ImageStore{a: 1, b: 1}

	ImageStoreLog = append(ImageStoreLog, imStore1)
	LogIndex++
	m[imStore1.a] = &ImageStoreLog[LogIndex]

	ans := *m[imStore1.a]
	fmt.Println(ans.b)

	ImageStoreLog = append(ImageStoreLog, imStore2)
	LogIndex++
	m[imStore2.a] = &ImageStoreLog[LogIndex]

	ImageStoreLog = append(ImageStoreLog, imStore3)
	LogIndex++
	m[imStore3.a] = &ImageStoreLog[LogIndex]

	c := &ans
	fmt.Println(c)
	fmt.Println(c++)
}
