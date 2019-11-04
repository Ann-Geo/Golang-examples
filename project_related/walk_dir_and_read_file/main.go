package main


import (
	"fmt"
	"os"
	"path/filepath"
)



func main() {

imageFilesPath := "/home/research/goworkspace/src/github.com/arun-ravindran/test_images/50K/"

fileList := walkAllFilesInDir(imageFilesPath)

fmt.Println(fileList)

}


func walkAllFilesInDir(dir string) []string {
	fileList := make([]string, 0)
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			fileList = append(fileList, path)
			fmt.Println(path)
		}
		return err
	})

	return fileList
}
