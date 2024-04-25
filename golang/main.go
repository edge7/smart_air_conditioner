package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	air "air_driver/interface_to_air"
	cam "air_driver/interface_to_camera"
	py "air_driver/interfaces_to_python"
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
	json.NewEncoder(w).Encode(acState)
}

// toggleHandler toggles the AC's power state
func toggleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	mutex.Lock()
	status, err := getCurrentStatus()
	if err != nil {
		log.Println("Error getting current status: ", err)
		status = "off"
	} else {
		if status == "on" {
			air.SendIRCommand("off")

		} else {
			air.SendIRCommand("on")
		}
		status, _ = getCurrentStatus()
	}
	if status == "on" {
		acState.IsOn = true
	} else {
		acState.IsOn = false
	}
	mutex.Unlock()

	json.NewEncoder(w).Encode(acState)
}

func getCurrentStatus() (string, error) {
	err := cam.TakePhoto()
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

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	err, temperature := py.GetTemperature()

	if err != nil {
		log.Printf("could not get model result: %v\n", err)
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

	time.Sleep(50000 * time.Second)
}
