package bell

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
	"time"
)

var (
	err     error
	btnPin = rpio.Pin(14)
	ledPin = rpio.Pin(15)
	numBlink = 8
	blinkTime = 0.3
	gongURL = "http://192.168.1.5/gong"
)

func StartBell(){
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		fmt.Printf("This is probably not running on a RaspberryPI\n")
		return
	}

//defer rpio.Close()
	ledPin.Output()

	btnPin.Input()
	btnPin.PullDown()
	btnPin.Detect(rpio.RiseEdge) // enable falling edge event detection

	go startBellWatcher()
}

func startBellWatcher() {
	for {
		if btnPin.EdgeDetected() { // check if event occured
			for x := 0; x < numBlink; x++ {
				ledPin.Toggle()
				time.Sleep(time.Second / 5)
			}
			resp, err := http.Get(gongURL)
			if err != nil {
				fmt.Printf("Could not find Gong\n")
			}
			resp.Body.Close()
		}
		time.Sleep(time.Second / 2)
		fmt.Printf("This is probably not running on a RaspberryPI\n")
	}
}
