package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Hello")
	defer window.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	classifier.Load("haarcascade_frontalface_default.xml")

	for {
		img := gocv.NewMat()
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device\n")
			return
		}
		if img.Empty() {
			continue
		}

		myface := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces", len(myface))

		for _, r := range myface {
			gocv.Rectangle(&img, r, color.RGBA{0, 255, 0, 0}, 3)
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
