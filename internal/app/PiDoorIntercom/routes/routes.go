package routes

import (
	"github.com/FloThinksPi/PiDoorIntercom/internal/app/PiDoorIntercom/camera"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func init() {
}

var mutex = &sync.Mutex{}

// Persist saves the given data
func Video(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	data := ""
	for {
		mutex.Lock()
		data = "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(camera.GetFrame()) + "\r\n\r\n"
		mutex.Unlock()
		time.Sleep(66 * time.Millisecond)
		_, _ = c.Writer.Write([]byte(data))
	}
}
