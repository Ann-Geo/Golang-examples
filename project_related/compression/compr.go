package main

import (
	//"fmt"
	"os"
	//"bufio"
	"io/ioutil"
	"bytes"
	"compress/gzip"
	"image/png"

)



func compress(input []byte) []byte {
        var buf bytes.Buffer
        compr := gzip.NewWriter(&buf)
        compr.Write(input) // here it appears to hang until
                           // Enter is pressed
        
        output := buf.Bytes()
	compr.Close()
        return output
}


func main() {



new_image := "image0000000.png"
file, _ := os.Open(new_image)

src, _ := png.Decode(file)

buf := new(bytes.Buffer)
_ = png.Encode(buf, src)
im_bytes := buf.Bytes()

compressed := compress(im_bytes)

ioutil.WriteFile("compressed_image.png", compressed, 0644)



/*
	fileToBeUploaded := "image0000000.png"
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
        	fmt.Println(err)
        	os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	im_bytes, _ := buffer.Read(bytes)    // <--------------- here!

	compressed := compress(im_bytes)

	ioutil.WriteFile("compressed_image.png", compressed, 0644)



	// convert []byte to image for saving to file
	img, _, _ := image.Decode(bytes.NewReader(compressed))

   	//save the imgByte to file
   	out, err := os.Create("QRImg.png")

   	if err != nil {
             fmt.Println(err)
             os.Exit(1)
   	}

   	err = png.Encode(out, img)

   	if err != nil {
            fmt.Println(err)
            os.Exit(1)
   	}
*/

}
