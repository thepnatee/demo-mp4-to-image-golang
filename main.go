package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	vidio "github.com/AlexEidt/Vidio"
)

func HTTPDownload(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ReadFile: Size of download: %d\n", len(d))
	return d, err
}

func WriteFile(dst string, d []byte) error {
	fmt.Printf("WriteFile: Size of download: %d\n", len(d))
	err := ioutil.WriteFile(dst, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func DownloadToFile(uri string, dst string) {
	if d, err := HTTPDownload(uri); err == nil {
		if WriteFile(dst, d) == nil {
			fmt.Printf("saved %s as %s\n", uri, dst)
		}
	}
}

func main() {

	// Image URL
	var imageUrl = "https://firebasestorage.googleapis.com/v0/b/line-lab-demo01.appspot.com/o/video.mp4?alt=media"

	// Get Current Dir
	path, _ := os.Getwd()

	// Down Load Image URL to File
	DownloadToFile(imageUrl, path+"/video.mp4")

	// Create File VDO
	video, _ := vidio.NewVideo("video.mp4")
	/*
		หากมี file ในเครื่องแล้วสามารถ Read จาก Folder ได้เลย
	*/

	img := image.NewRGBA(image.Rect(0, 0, video.Width(), video.Height()))
	video.SetFrameBuffer(img.Pix)

	frame := 0
	for video.Read() {
		f, _ := os.Create(fmt.Sprintf("img/%d.jpg", frame))
		jpeg.Encode(f, img, nil)
		f.Close()
		// frame++ save image จนจบเพลง
	}
}
