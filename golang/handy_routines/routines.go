package handy_routines

import (
	cam "air_driver/interface_to_camera"
	py "air_driver/interfaces_to_python"
)

func GetTemperatureGoRoutine(result chan float32, errorChan chan error) {
	err, temperature := py.GetTemperature()
	if err != nil {
		errorChan <- err
	} else {
		result <- temperature
	}

}

func GetStatusGoRoutine(result chan string, errorChan chan error) {
	status, err := cam.GetCurrentStatus()
	if err != nil {
		errorChan <- err
	} else {
		result <- status
	}
}
