package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read image from file that already exists
	existingImageFile, err := ioutil.ReadFile("/home/research/goworkspace/src/vsc_workspace/images/image0000000.png")
	if err != nil {
		// Handle error
	}
	//defer existingImageFile.Close()

	fmt.Printf("read format %T\n", existingImageFile)

	//toBytes := existingImageFile.WriteString

	//fmt.Printf("byte format %T\n", toBytes)

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	//imageData, imageType, err := image.Decode(existingImageFile)
	//if err != nil {
	// Handle error
	//}

	//fmt.Printf("image data format %T\n", imageData)
	//fmt.Println(imageType)

	//fileInfo, _ := existingImageFile.Stat()
	//var size int64 = fileInfo.Size()
	//bytes := make([]byte, size)

	//buffer := bufio.NewReader(existingImageFile)
	//toBytes1, err := buffer.Read(bytes)

	//fmt.Printf("buffer format %T\n", toBytes1)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	//existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	//_, err1 := png.Decode(existingImageFile)
	//if err1 != nil {
	// Handle error
	//}
	//fmt.Println(loadedImage)
}
