package main

import (
	"gocv.io/x/gocv"
	"log"
	"image/color"
	"fmt"
)

func main() {
	deviceID := 0

	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		log.Fatalf("Could not open capture device %v", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	blue := color.RGBA{0, 0, 255, 0}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("data/haarcascade_frontalface_default.xml")

	fmt.Printf("Start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image
		for _, r := range rects {
			gocv.Rectangle(img, r, blue, 3)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)
	}

}
