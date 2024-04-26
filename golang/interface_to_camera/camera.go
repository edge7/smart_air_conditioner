package interface_to_camera

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

var photoMutex sync.Mutex

func TakePhoto() error {
	photoMutex.Lock()
	defer photoMutex.Unlock()
	currentHour := time.Now().Hour()
	gain := getGain(currentHour)
	log.Println("current hour:", currentHour)
	log.Println("gain:", gain)

	cmd := exec.Command(
		"libcamera-still",
		"--width", "512",
		"--height", "512",
		"--shutter", "1000000",
		"--gain", fmt.Sprint(gain),
		"--denoise", "cdn_fast",
		"--output", "/tmp/fresh_image.jpg",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("libcamera command failed with", err)
		return fmt.Errorf("command failed with %s: %s", err, output)
	}
	log.Println("libcamera command executed successfully")
	//log.Println("output:", string(output))
	return nil
}

func getGain(hour int) int {
	switch {
	case hour < 5:
		return 4
	case hour >= 5 && hour < 9:
		return 3
	case hour >= 9 && hour < 19:
		return 0
	case hour == 19 || hour == 20:
		return 2
	case hour > 20:
		return 5
	default:
		return 0 // Default case, should not hit due to complete coverage above
	}
}
