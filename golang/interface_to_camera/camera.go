package interface_to_camera

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	py "air_driver/interfaces_to_python"
)

var photoMutex sync.Mutex

const ImgPath = "/tmp/fresh_image.jpg"

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
		"--output", ImgPath,
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

func GetCurrentStatus(takePicture bool) (string, error) {
	var err error = nil
	if takePicture {
		err = TakePhoto()
	} else {
		log.Println("Not taking a picture, using cached image")
	}
	if err != nil {
		log.Println("Error taking photo: ", err)
		return "", err
	} else {
		err, modelPred := py.GetModelResult("")
		if err != nil {
			return "", err
		}
		return modelPred, nil
	}
}
