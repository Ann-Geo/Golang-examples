package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func run() ([]string, error) {
	searchDir := "/home/research/goworkspace/src/vsc_workspace/images"

	fileList := make([]string, 0)
	fmt.Println(fileList)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		//fmt.Println(fileList)
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil
}

func main() {
	run()
}
