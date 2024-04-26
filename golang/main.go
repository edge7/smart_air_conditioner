package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	handy_routine "air_driver/handy_routines"
	air "air_driver/interface_to_air"
	cam "air_driver/interface_to_camera"
	py "air_driver/interfaces_to_python"
	notif "air_driver/notifications"
)

type ACState struct {
	IsOn bool `json:"isOn"`
}

// Global AC state variable with a mutex for safe access across goroutines
var acState ACState
var mutex sync.Mutex

func statusHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	log.Println("Status requested")
	status, _ := cam.GetCurrentStatus()
	log.Println("status is ", status)
	if status == "on" {
		acState.IsOn = true
	} else {
		acState.IsOn = false
	}
	json.NewEncoder(w).Encode(acState)
}

// toggleHandler toggles the AC's power state
func toggleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Toggle requested")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	mutex.Lock()
	status, err := cam.GetCurrentStatus()
	log.Println("asking for current status:", status)
	if err != nil {
		log.Println("Error getting current status: ", err)
		status = "off"
	} else {
		if status == "on" {
			air.SendIRCommand("off")

		} else {
			air.SendIRCommand("on")
		}
		status, _ = cam.GetCurrentStatus()
	}
	if status == "on" {
		acState.IsOn = true
	} else {
		acState.IsOn = false
	}
	mutex.Unlock()

	json.NewEncoder(w).Encode(acState)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Asking Temperature")
	err, temperature := py.GetTemperature()

	if err != nil {
		log.Printf("could not get Temperature result: %v\n", err)
		temperature = -1.0
	}

	log.Println("Temperature:", temperature)
	tempData := struct {
		Temperature float64 `json:"temperature"`
	}{
		Temperature: float64(temperature),
	}

	json.NewEncoder(w).Encode(tempData)
}

// imageHandler serves a static image URL
func imageHandler(w http.ResponseWriter, r *http.Request) {
	err := cam.TakePhoto()
	if err != nil {
		log.Println("Failed to take photo: %v", err)
	}
	imagePath := "/tmp/fresh_image.jpg"
	http.ServeFile(w, r, imagePath)
}

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Should be restricted in production
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set(
			"Access-Control-Allow-Headers", "Content-Type, Authorization",
		) // Add other headers as needed

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}
func main() {

	go func() {
		http.HandleFunc("/status", cors(statusHandler))
		http.HandleFunc("/toggle", cors(toggleHandler))
		http.HandleFunc("/image", cors(imageHandler))
		http.HandleFunc("/temperature", cors(temperatureHandler))
		http.Handle(
			"/", http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// Check if file exists and, if not, use index.html
					path := "./build" + r.URL.Path
					log.Println("Path:", path)
					if _, err := os.Stat(path); os.IsNotExist(err) {
						http.ServeFile(w, r, "./build/index.html")
					} else {
						http.ServeFile(w, r, path)
					}
				},
			),
		)

		log.Println("Server starting on port 3030...")
		if err := http.ListenAndServe(":3030", nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	log.Println("Server started on port 3030, now handling business logic...")
	notif.SendNotification("Starting APP", "WHARFREE")
	for {
		var temperature float32
		var errTemp error = nil
		var errStatus error = nil
		var status string
		const OFF = "off"
		const ON = "on"

		tempChan := make(chan float32)
		statusChan := make(chan string)
		errChanTemp := make(chan error)
		errChanStatus := make(chan error)
		go handy_routine.GetTemperatureGoRoutine(tempChan, errChanTemp)
		go handy_routine.GetStatusGoRoutine(statusChan, errChanStatus)

		for i := 0; i < 2; i++ { // We expect 2 responses
			select {
			case temp := <-tempChan:
				temperature = temp
			case stat := <-statusChan:
				status = stat
			case e := <-errChanTemp:
				errTemp = e
				log.Printf(
					"Temperature Thread: Could not get Temperature result: %v\n",
					errTemp,
				)
			case e := <-errChanStatus:
				errStatus = e
				log.Printf("Status Thread: Could not get Status result: %v\n", errStatus)
			}
		}

		// Here we have the temperature and the status and potential errors

		// If the temperature is not valid, we will use the time to act proactively.
		if errTemp != nil {
			log.Printf("Checking the time to act proactively")
			currentHour := time.Now().Hour()
			log.Println("current hour is:", currentHour)
			if currentHour >= 9 && currentHour < 20 {
				log.Println("It is between 9am and 8pm, so we will turn on the AC, in any case")
				temperature = 30
			} else {
				log.Println("It is not between 9am and 8pm, so we will not turn on the AC")
				temperature = 10
			}
		}

		log.Println("Temperature is:", temperature)

		if temperature > 24 {

			// Here it's hot

			if errStatus != nil {
				log.Println("Error getting current status: ", errStatus)
				status = OFF
			}
			if status == OFF {
				maxAttempts := 3
				turnXReliable(maxAttempts, ON)

			} else {
				log.Println("Status is on and need to stay on.. so nothing to do")
			}
		} else {

			// Here it's NOT hot
			log.Println("Need to turn off the AC")

			if errStatus != nil {
				log.Println("Error getting current status: ", errStatus)
				status = ON
			}
			if status == ON {

				maxAttempts := 3
				turnXReliable(maxAttempts, OFF)

			} else {
				log.Println("Status is off and need to stay off.. so nothing to do")
			}
		}
		time.Sleep(60 * 10 * time.Second)

	}

}

func turnXReliable(
	maxAttempts int, doWhat string,
) {
	var errStatus error
	var status string
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		errIr := air.SendIRCommand(doWhat)
		if errIr != nil {
			log.Println("Error sending IR command: ", errIr)
			if attempt == maxAttempts {
				log.Printf("Attempt %d failed to sendIR: %v\n", attempt, errIr)
				notif.SendNotification("Failed to turn AC "+doWhat, "WHARFREE")
			}
		} else {
			// Need to take a new photo before checking
			status, errStatus = cam.GetCurrentStatus()
			if errStatus != nil {
				log.Println("Error getting current status: ", errStatus)
				status = ""
			}
			if status == doWhat {
				log.Println("AC turned on successfully")
				notif.SendNotification("AC turned successfully "+doWhat, "WHARFREE")
				break
			}
		}
		time.Sleep(10 * time.Second)

	}
}
