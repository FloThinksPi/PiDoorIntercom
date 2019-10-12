package camera

import (
	"fmt"
	"gocv.io/x/gocv"
)

var (
	err     error
	camera  *gocv.VideoCapture
	frameId int
)

var frame []byte

func init() {

	camera, err = gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Printf("Error opening capture device: \n")
		return
	}

	go startStream()

}

func GetFrame() []byte {
	return frame
}

func startStream() {
	img := gocv.NewMat()
	defer img.Close()
	for {
		if ok := camera.Read(&img); !ok {
			fmt.Printf("Device closed\n")
			panic("Camera closed unexpectedly.")
		}
		if img.Empty() {
			continue
		}
		frameId++
		//gocv.Resize(img, &img, image.Point{}, float64(2), float64(2), 0)
		frame, _ = gocv.IMEncode(".jpg", img)
	}
}