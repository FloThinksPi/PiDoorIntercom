package bell

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
	"time"
)

var (
	err     error
	blinkTime = 300
	blinkCycles = 8
	gongURL = "http://192.168.1.5/gong"

)

func InitBellWatcher(){

	// Setup
	if rpio.Open() != nil {
		fmt.Printf("Error accesing GPIOs. This is likely not running on a RaspberryPI with Raspbian\n")
		return
	}
	btn := rpio.Pin(14)
	btn.Input()
	btn.PullDown()
	led := rpio.Pin(15)
	led.Output()

	// Bell Logic
	go pollBellButton(led,btn)

}

func pollBellButton(ledPin rpio.Pin,btnPin rpio.Pin){
	for true {
		if btnPin.Read() == rpio.High {
			print("Button Press detected\n")
			for p := 0; p < 2; p++ {
				_, err = http.Get(gongURL)
				for i := 0; i < blinkCycles; i++ {
					ledPin.High()
					time.Sleep(time.Millisecond * time.Duration(blinkTime))
					ledPin.Low()
					time.Sleep(time.Millisecond * time.Duration(blinkTime))
				}
			}
		}
		time.Sleep(time.Millisecond * 50)
	}
}